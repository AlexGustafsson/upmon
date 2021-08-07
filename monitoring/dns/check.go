package dns

import (
	"net"
	"sync"
	"time"

	"github.com/AlexGustafsson/upmon/monitoring/core"
	log "github.com/sirupsen/logrus"
)

type Check struct {
	config *MonitorConfiguration
}

// NewCheck creates a new check for a service
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
	return "dns"
}

func (check *Check) Description() string {
	return "DNS resolver"
}

func (check *Check) Version() string {
	return "0.1.0"
}

func (check *Check) Perform() (core.Status, error) {
	_, err := net.LookupIP(check.config.Hostname)
	if err != nil {
		return core.StatusUnknown, err
	}

	return core.StatusUp, nil
}

func (check *Check) Config() core.MonitorConfiguration {
	return check.config
}

func (check *Check) Watch(update chan<- *core.ServiceStatus, stop <-chan bool, wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-stop:
				wg.Done()
				log.WithFields(log.Fields{"check": "dns"}).Debugf("exiting")
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
