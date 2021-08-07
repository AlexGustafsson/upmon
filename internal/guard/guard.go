package guard

import (
	"fmt"
	"sync"

	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/monitoring"
	log "github.com/sirupsen/logrus"
)

// Guard is a monitoring manager
type Guard struct {
	sync.Mutex
	configuredServices []*configuration.ServiceConfiguration
	configuredMonitors []*monitoring.Monitor
	activeMonitors     []*monitoring.Monitor
	StatusUpdates      chan *monitoring.MonitorStatus
}

func NewGuard() *Guard {
	guard := &Guard{
		StatusUpdates:      make(chan *monitoring.MonitorStatus),
		configuredServices: make([]*configuration.ServiceConfiguration, 0),
		configuredMonitors: make([]*monitoring.Monitor, 0),
		activeMonitors:     make([]*monitoring.Monitor, 0),
	}

	return guard
}

// ConfigureServices sets up the configured monitors. Use Reload to apply the configuration
func (guard *Guard) ConfigureServices(services []*configuration.ServiceConfiguration) error {
	guard.Lock()
	defer guard.Unlock()

	guard.configuredServices = services
	guard.configuredMonitors = guard.configuredMonitors[:0]

	// Create all configured monitors
	// TODO: Transaction / rollback - either all monitors start or none start (no undefined state)
	for _, serviceConfig := range guard.configuredServices {
		service := &monitoring.Service{
			Id:          serviceConfig.Id,
			Name:        serviceConfig.Name,
			Description: serviceConfig.Description,
			Origin:      serviceConfig.Origin,
		}

		for _, monitorConfig := range serviceConfig.Monitors {
			check, err := monitoring.NewCheck(monitorConfig.Type, monitorConfig.Options)
			if err != nil {
				log.Warningf("failed to create monitor '%s' (%s) for service '%s' (%s): %v", monitorConfig.Name, monitorConfig.Id, serviceConfig.Name, serviceConfig.Id, err)
				continue
			}

			errs := check.Config().Validate()
			if len(errs) != 0 {
				for _, err := range errs {
					log.Error(err)
				}
				return fmt.Errorf("monitor config validation failed for service '%s' (%s), monitor '%s' (%s)", serviceConfig.Name, serviceConfig.Id, monitorConfig.Name, monitorConfig.Id)
			}

			monitor := &monitoring.Monitor{
				Id:          monitorConfig.Id,
				Name:        monitorConfig.Name,
				Description: monitorConfig.Description,
				Check:       check,
				Service:     service,
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

	log.Infof("starting all monitors")
	for _, monitor := range guard.configuredMonitors {
		log.Infof("starting monitor '%s' (%s) for service '%s' (%s)", monitor.Name, monitor.Id, monitor.Service.Name, monitor.Service.Id)
		err := monitor.Watch(guard.StatusUpdates)
		if err != nil {
			log.Warningf("failed to start watching monitor '%s' (%s) for service '%s' (%s): %v", monitor.Name, monitor.Id, monitor.Service.Name, monitor.Service.Id, err)
		}
		guard.activeMonitors = append(guard.activeMonitors, monitor)
	}
}

// stopAllMonitors stops all monitors and waits for them to close
func (guard *Guard) stopAllMonitors() {
	guard.Lock()
	defer guard.Unlock()

	log.Infof("stopping all monitors")
	var wg sync.WaitGroup
	for _, monitor := range guard.activeMonitors {
		go monitor.Stop(&wg)
	}
	wg.Wait()
	guard.activeMonitors = guard.activeMonitors[:0]
	log.Infof("all monitors stopped successfully")
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
