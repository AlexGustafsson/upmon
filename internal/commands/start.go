package commands

import (
	"encoding/json"
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
	self     string
	welcomer func(node *memberlist.Node)
}

type delegate struct{}

func (delegate *delegate) NodeMeta(limit int) []byte {
	bytes := make([]byte, 0)
	return bytes
}

func (delegate *delegate) NotifyMsg(message []byte) {
	// TODO: apply config of distributed services
	log.Info("received message: %v", string(message))
}

func (delegate *delegate) GetBroadcasts(overhead, limit int) [][]byte {
	broadcasts := make([][]byte, 0)
	return broadcasts
}

func (delegate *delegate) LocalState(join bool) []byte {
	bytes := make([]byte, 0)
	return bytes
}

func (delegate *delegate) MergeRemoteState(buf []byte, join bool) {

}

func (delegate *eventDelegate) NotifyJoin(node *memberlist.Node) {
	log.WithFields(log.Fields{"name": node.Name, "address": node.Addr, "port": node.Port}).Info("node joined")
	if node.Name != delegate.self {
		if delegate.welcomer != nil {
			delegate.welcomer(node)
		}
	}
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

	memberlistConfig, err := config.MemberlistConfig()
	if err != nil {
		return err
	}
	eventDelegate := &eventDelegate{
		self: config.Name,
	}
	memberlistConfig.Events = eventDelegate
	memberlistConfig.Conflict = &conflictDelegate{}
	memberlistConfig.Delegate = &delegate{}

	log.WithFields(log.Fields{"bind": config.Bind}).Info("listening")

	list, err := memberlist.Create(memberlistConfig)
	if err != nil {
		return err
	}

	eventDelegate.welcomer = func(node *memberlist.Node) {
		log.WithFields(log.Fields{"name": node.Name}).Info("welcoming node")
		message, err := json.Marshal(config)
		if err != nil {
			log.WithFields(log.Fields{"name": node.Name}).Errorf("failed to welcome node: %v", err)
			return
		}

		err = list.SendReliable(node, message)
		if err != nil {
			log.WithFields(log.Fields{"name": node.Name}).Errorf("failed to welcome node: %v", err)
			return
		}
	}

	if len(config.Peers) > 0 {
		var delay time.Duration = time.Second
		var connectionError error
		for i := 0; i < 5; i++ {
			log.Infof("attempting to join cluster (attempt %d/5)", i+1)
			contactedPeers, err := list.Join(config.Peers)
			if err == nil {
				log.Infof("made contact with %d peers", contactedPeers)
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
		go server.Start(config.Api.Bind)
	}

	wg.Wait()

	return nil
}
