package guard

import (
	"fmt"
	"sync"

	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/monitor"
	"github.com/AlexGustafsson/upmon/monitor/core"
	log "github.com/sirupsen/logrus"
)

// Guard is a monitoring manager
type Guard struct {
	sync.Mutex
	monitorsStop       chan bool
	monitorsGroup      sync.WaitGroup
	configuredServices []configuration.ServiceConfiguration
	configuredMonitors []core.Monitor
	activeMonitors     []core.Monitor
	StatusUpdates      chan *core.ServiceStatus
}

func NewGuard() *Guard {
	guard := &Guard{
		StatusUpdates:      make(chan *core.ServiceStatus),
		configuredServices: make([]configuration.ServiceConfiguration, 0),
		configuredMonitors: make([]core.Monitor, 0),
		activeMonitors:     make([]core.Monitor, 0),
	}

	return guard
}

// ConfigureServices sets up the configured monitors. Use Reload to apply the configuration
func (guard *Guard) ConfigureServices(services []configuration.ServiceConfiguration) error {
	guard.Lock()
	defer guard.Unlock()

	guard.configuredServices = services
	guard.configuredMonitors = guard.configuredMonitors[:0]

	// Create all configured monitors
	// TODO: Transaction / rollback - either all monitors start or none start (no undefined state)
	for _, serviceConfig := range guard.configuredServices {
		service := core.NewService(serviceConfig.Id, serviceConfig.Name, serviceConfig.Description)

		for _, monitorConfig := range serviceConfig.Monitors {
			monitor, err := monitor.NewMonitor(monitorConfig.Type, service, monitorConfig.Options)
			if err != nil {
				log.Warningf("failed to create monitor '%s' (%s) for service '%s': %v", monitorConfig.Name, monitorConfig.Type, serviceConfig.Name, err)
				continue
			}

			errs := monitor.Config().Validate()
			if len(errs) != 0 {
				for _, err := range errs {
					log.Error(err)
				}
				return fmt.Errorf("monitor config validation failed for service '%s', monitor '%s' (%s)", serviceConfig.Name, monitorConfig.Name, monitorConfig.Type)
			}

			guard.configuredMonitors = append(guard.configuredMonitors, monitor)
		}
	}

	return nil
}

// Start starts the guard
func (guard *Guard) Start() {
	guard.startAllMonitors()
}

// Stop stops the guard
func (guard *Guard) Stop() {
	guard.stopAllMonitors()
}

// startAllMonitors starts all monitors. Calling when running is undefined behavior
func (guard *Guard) startAllMonitors() {
	guard.Lock()
	defer guard.Unlock()

	// TODO: Rollback if all monitors can't start? Make atomic?
	guard.activeMonitors = guard.activeMonitors[:0]
	guard.monitorsStop = make(chan bool)

	log.Infof("starting all monitors")
	for _, monitor := range guard.configuredMonitors {
		log.Infof("starting monitor '%s' for service '%s'", monitor.Name(), monitor.Service().Name())
		err := monitor.Watch(guard.StatusUpdates, guard.monitorsStop, guard.monitorsGroup)
		if err != nil {
			log.Warningf("failed to start watching '%s' (%s): %v", monitor.Name(), monitor.Type(), err)
		}
		guard.activeMonitors = append(guard.activeMonitors, monitor)
	}
	log.Infof("all monitors have started")
}

// stopAllMonitors stops all monitors and waits for them to close
func (guard *Guard) stopAllMonitors() {
	if guard.monitorsStop == nil {
		return
	}

	guard.Lock()
	defer guard.Unlock()

	log.Infof("stopping all monitors")
	close(guard.monitorsStop)
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
