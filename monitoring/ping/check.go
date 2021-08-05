package ping

import (
	"sync"
	"time"

	"github.com/AlexGustafsson/upmon/monitoring/core"
	"github.com/go-ping/ping"
	log "github.com/sirupsen/logrus"
)

type Check struct {
	config *MonitorConfiguration
}

// NewCheck creates a new monitor for a service
func NewCheck(options map[string]interface{}) (*Check, error) {
	config, err := ParseConfiguration(options)
	if err != nil {
		return nil, err
	}

	return &Check{
		config: config,
	}, nil
}

func (check *Check) Name() string {
	return "ping"
}

func (check *Check) Description() string {
	return "Pinging check"
}

func (check *Check) Version() string {
	return "0.1.0"
}

func (check *Check) Config() core.MonitorConfiguration {
	return check.config
}

func (check *Check) Perform() (core.Status, error) {
	pinger := ping.New(check.config.Hostname)
	pinger.Count = check.config.Count
	pinger.Timeout = check.config.Timeout

	log.WithFields(log.Fields{"check": "ping"}).Debugf("resolving '%s'", check.config.Hostname)
	err := pinger.Resolve()
	if err != nil {
		return core.StatusUnknown, err
	}

	log.WithFields(log.Fields{"check": "ping"}).Debugf("pinging %s (%s)", pinger.IPAddr(), check.config.Hostname)
	err = pinger.Run()
	if err != nil {
		return core.StatusUnknown, err
	}

	statistics := pinger.Statistics()
	log.WithFields(log.Fields{"check": "ping"}).Debugf("sent %d packets, received %d packets, loss is %.0f%%", statistics.PacketsSent, statistics.PacketsRecv, statistics.PacketLoss)
	if statistics.PacketLoss == 0 {
		return core.StatusUp, nil
	}

	return core.StatusDown, nil
}

func (check *Check) Watch(update chan<- *core.ServiceStatus, stop <-chan bool, wg sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-stop:
				wg.Done()
				log.WithFields(log.Fields{"check": "ping"}).Debugf("exiting")
				return
			default:
				status, err := check.Perform()
				update <- &core.ServiceStatus{
					Err:    err,
					Status: status,
					Check:  check,
				}
			}
			time.Sleep(check.config.Interval)
		}
	}()

	return nil
}
