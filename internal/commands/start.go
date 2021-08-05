package commands

import (
	"fmt"
	"sync"
	"time"

	"github.com/AlexGustafsson/upmon/internal/clustering"
	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/internal/guard"
	"github.com/AlexGustafsson/upmon/internal/server"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type customFormatter struct {
	nodeName  string
	formatter log.Formatter
}

func (formatter customFormatter) Format(entry *log.Entry) ([]byte, error) {
	entry.Data["node"] = formatter.nodeName
	return formatter.formatter.Format(entry)
}

func startCommand(context *cli.Context) error {
	configPath := context.String("config")

	config, err := configuration.Load(configPath)
	if err != nil {
		return err
	}

	errors := configuration.Validate(config)
	if len(errors) != 0 {
		for _, err := range errors {
			log.Error(err)
		}
		return fmt.Errorf("failed to validate the configuration")
	}

	log.SetFormatter(&customFormatter{
		nodeName: config.Name,
		formatter: &log.TextFormatter{
			ForceQuote:       true,
			FullTimestamp:    true,
			QuoteEmptyFields: true,
		},
	})

	cluster, err := clustering.NewCluster(config)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"bind": config.Bind}).Info("listening")

	if len(config.Peers) > 0 {
		var delay time.Duration = time.Second
		var connectionError error
		for i := 0; i < 5; i++ {
			log.Infof("attempting to join cluster (attempt %d/5)", i+1)
			err := cluster.Join(config.Peers)
			if err == nil {
				connectionError = nil
				break
			}

			log.Errorf("failed to join cluster, trying again in %.0fs", delay.Seconds())
			time.Sleep(delay)
			delay *= 2
			connectionError = err
		}
		if connectionError != nil {
			return connectionError
		}

		health := cluster.Status()
		log.Infof("cluster status is %s", health.String())
	} else {
		log.Warning("no peers configured")
	}

	guard := guard.NewGuard()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for update := range cluster.ServicesUpdates {
			log.Debug("got service update from cluster")
			guard.ConfigureServices(update)
			guard.Reload()
		}
	}()

	wg.Add(1)
	go guard.Start()

	if config.Api.Enabled {
		server := server.NewServer(config, cluster)
		wg.Add(1)
		go server.Start(config.Api.Bind)
	}

	wg.Wait()

	return nil
}
