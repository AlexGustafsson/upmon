package guard

import (
	"fmt"
	"sync"

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
	stop        chan bool
}

// Service is a monitored service
type Service struct {
	name        string
	description string
}

// Guard is a monitoring manager
type Guard struct {
	sync.Mutex
	monitorsGroup      sync.WaitGroup
	update             chan *core.ServiceStatus
	stop               chan bool
	configuredServices map[string]configuration.ServiceConfiguration
	configuredMonitors []*Monitor
	activeMonitors     []*Monitor
}

func NewGuard() *Guard {
	guard := &Guard{
		update:             make(chan *core.ServiceStatus),
		configuredServices: make(map[string]configuration.ServiceConfiguration),
		configuredMonitors: make([]*Monitor, 0),
		activeMonitors:     make([]*Monitor, 0),
	}

	return guard
}

// ConfigureServices sets up the configured monitors. Use Reload to apply the configuration
func (guard *Guard) ConfigureServices(services map[string]configuration.ServiceConfiguration) error {
	guard.Lock()
	defer guard.Unlock()

	guard.configuredServices = services
	guard.configuredMonitors = guard.configuredMonitors[:0]

	// Create all configured monitors
	// TODO: Transaction / rollback - either all monitors start or none start (no undefined state)
	for serviceName, serviceConfig := range guard.configuredServices {
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
				return fmt.Errorf("monitor config validation failed for service '%s', monitor '%s' (%s)", serviceName, monitorConfig.Name, monitorConfig.Type)
			}

			guard.configuredMonitors = append(guard.configuredMonitors, &Monitor{
				name:        monitorConfig.Name,
				description: monitorConfig.Description,
				monitor:     monitor,
				stop:        make(chan bool),
			})
		}
	}

	return nil
}

// Start starts the guard
func (guard *Guard) Start() error {
	guard.startAllMonitors()
	guard.stop = make(chan bool)

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

func (guard *Guard) Stop() {
	close(guard.stop)
}

// startAllMonitors starts all monitors. Calling when running is undefined behavior
func (guard *Guard) startAllMonitors() {
	guard.Lock()
	defer guard.Unlock()

	// TODO: Rollback if all monitors can't start? Make atomic?
	guard.activeMonitors = guard.activeMonitors[:0]

	log.Infof("starting all monitors")
	for _, monitor := range guard.configuredMonitors {
		log.Infof("starting monitor '%s'", monitor.name)
		err := monitor.monitor.Watch(guard.update, monitor.stop, guard.monitorsGroup)
		if err != nil {
			log.Warningf("failed to start watching '%s' (%s): %v", monitor.name, monitor.monitor.Name(), err)
		}
		guard.activeMonitors = append(guard.activeMonitors, monitor)
	}
	log.Infof("all monitors have started")
}

// stopAllMonitors stops all monitors and waits for them to close
func (guard *Guard) stopAllMonitors() {
	guard.Lock()
	defer guard.Unlock()

	log.Infof("stopping all monitors")
	for _, monitor := range guard.activeMonitors {
		log.Infof("stopping monitor '%s'", monitor.name)
		close(monitor.stop)
	}
	log.Infof("waiting for all monitors to stop")
	guard.monitorsGroup.Wait()
	guard.activeMonitors = guard.activeMonitors[:0]
	log.Infof("all monitors have stopped")
}

func (guard *Guard) Reload() error {
	guard.stopAllMonitors()

	err := guard.ConfigureServices(guard.configuredServices)
	if err != nil {
		return err
	}

	guard.startAllMonitors()

	return nil
}
