package guard

import (
	"fmt"

	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/monitor"
	"github.com/AlexGustafsson/upmon/monitor/core"
	log "github.com/sirupsen/logrus"
)

// Monitor is a monitor of a service
type Monitor struct {
	name        string
	description string
	monitor     core.Monitor
	service     *Service
	stop        <-chan bool
}

// Service is a monitored service
type Service struct {
	name        string
	description string
}

// Guard is a monitoring manager
type Guard struct {
	config   *configuration.Configuration
	monitors []*Monitor
	update   chan *core.ServiceStatus
	stop     <-chan bool
}

func NewGuard(config *configuration.Configuration) (*Guard, error) {
	guard := &Guard{
		config:   config,
		monitors: make([]*Monitor, 0),
		update:   make(chan *core.ServiceStatus),
	}

	// Create all configured monitors
	for serviceName, serviceConfig := range guard.config.Services {
		service := &Service{
			name:        serviceName,
			description: serviceConfig.Description,
		}

		for _, monitorConfig := range serviceConfig.Monitors {
			monitor, err := monitor.NewMonitor(monitorConfig.Type, service, monitorConfig.Options)
			if err != nil {
				log.Warningf("failed to create monitor '%s' (%s): %v", monitorConfig.Name, monitorConfig.Type, err)
				continue
			}

			errs := monitor.Config().Validate()
			if len(errs) != 0 {
				for _, err := range errs {
					log.Error(err)
				}
				return nil, fmt.Errorf("monitor config validation failed for service '%s', monitor '%s' (%s)", serviceName, monitorConfig.Name, monitorConfig.Type)
			}

			guard.monitors = append(guard.monitors, &Monitor{
				name:        monitorConfig.Name,
				description: monitorConfig.Description,
				monitor:     monitor,
				stop:        make(<-chan bool),
			})
		}
	}

	return guard, nil
}

// Start starts the guard
func (guard *Guard) Start() error {
	// Start all monitors
	for _, monitor := range guard.monitors {
		err := monitor.monitor.Watch(guard.update, monitor.stop)
		if err != nil {
			log.Warningf("failed to start watching '%s' (%s): %v", monitor.name, monitor.monitor.Name(), err)
		}
	}

	// Watch the update channel
	for {
		select {
		case <-guard.stop:
			return nil
		case status := <-guard.update:
			if status.Err == nil {
				log.Infof("got update: %s", status.Status.String())
			} else {
				log.Warningf("failed to perform '%s' check: %v", status.Monitor.Name(), status.Err)
			}
		}
	}
}
