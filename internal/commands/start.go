package commands

import (
	"fmt"
	"sync"
	"time"

	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/internal/guard"
	"github.com/AlexGustafsson/upmon/internal/server"
	"github.com/hashicorp/memberlist"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type eventDelegate struct {
}

func (delegate *eventDelegate) NotifyJoin(node *memberlist.Node) {
	log.WithFields(log.Fields{"name": node.Name, "address": node.Addr, "port": node.Port}).Info("node joined")
}

func (delegate *eventDelegate) NotifyLeave(node *memberlist.Node) {
	log.WithFields(log.Fields{"name": node.Name, "address": node.Addr, "port": node.Port}).Info("node left")
}

func (delegate *eventDelegate) NotifyUpdate(node *memberlist.Node) {
	log.WithFields(log.Fields{"name": node.Name, "address": node.Addr, "port": node.Port}).Info("node updated")
}

type conflictDelegate struct {
}

func (delegate *conflictDelegate) NotifyConflict(existing, other *memberlist.Node) {
	log.WithFields(log.Fields{"name": other.Name, "address": other.Addr, "port": other.Port}).Warning("name conflict")
}

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

	memberlistConfig := config.MemberlistConfig()
	memberlistConfig.Events = &eventDelegate{}
	memberlistConfig.Conflict = &conflictDelegate{}

	log.WithFields(log.Fields{"address": config.Address, "port": config.Port}).Info("listening")

	list, err := memberlist.Create(memberlistConfig)
	if err != nil {
		return err
	}

	if len(config.Peers) > 0 {
		var delay time.Duration = time.Second
		var connectionError error
		for i := 0; i < 5; i++ {
			log.Infof("attempting to join cluster (attempt %d/5)", i+1)
			contactedPeers, err := list.Join(config.PeerAddresses())
			if err == nil {
				log.Infof("made contact with %d peers", contactedPeers)
				err = nil
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

		health := list.GetHealthScore()
		healthDescription := ""
		switch health {
		case 0:
			healthDescription = "healthy"
		default:
			healthDescription = "not healthy"
		}
		log.Infof("health is at %d (%s)", health, healthDescription)
	} else {
		log.Warning("no peers configured")
	}

	guard, err := guard.NewGuard(config)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go guard.Start()

	if config.Api.Enabled {
		server := server.NewServer(config, list)
		wg.Add(1)
		go server.Start(config.Api.Address, config.Api.Port)
	}

	wg.Wait()

	return nil
}
