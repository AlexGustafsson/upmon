package commands

import (
	"fmt"
	"sync"

	"github.com/AlexGustafsson/upmon/internal/configuration"
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
		log.Info("attempting to join cluster")
		contacts, err := list.Join(config.PeerAddresses())
		if err != nil {
			return err
		}

		log.Infof("made contact with %d peers", contacts)

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

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	return nil
}
