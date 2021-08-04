package guard

import (
	"fmt"
	"time"

	"github.com/AlexGustafsson/upmon/internal/clustering"
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
	cluster  *clustering.Cluster
	monitors []*Monitor
	update   chan *core.ServiceStatus
	stop     <-chan bool
}

func NewGuard(cluster *clustering.Cluster) (*Guard, error) {
	guard := &Guard{
		cluster:  cluster,
		monitors: make([]*Monitor, 0),
		update:   make(chan *core.ServiceStatus),
	}

	err := guard.setupMonitors()
	if err != nil {
		return nil, err
	}

	return guard, nil
}

func (guard *Guard) setupMonitors() error {
	// Create all configured monitors
	for serviceName, serviceConfig := range guard.cluster.Services() {
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

			guard.monitors = append(guard.monitors, &Monitor{
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

	// TODO: Replace with an actual event bus for the cluster updates
	go func() {
		for range time.Tick(time.Second * 10) {
			log.Info("reloading config")
			err := guard.Reload()
			if err == nil {
				log.Info("reloaded successfully")
			} else {
				log.Errorf("unable to reload guard: %v", err)
			}
		}
	}()

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

func (guard *Guard) startAllMonitors() {
	// Start all monitors
	for _, monitor := range guard.monitors {
		log.Infof("starting monitor '%s'", monitor.name)
		err := monitor.monitor.Watch(guard.update, monitor.stop)
		if err != nil {
			log.Warningf("failed to start watching '%s' (%s): %v", monitor.name, monitor.monitor.Name(), err)
		}
	}
}

func (guard *Guard) stopAllMonitors() {
	for _, monitor := range guard.monitors {
		log.Infof("stopping monitor 's'", monitor.name)
		close(monitor.stop)
	}
}

func (guard *Guard) Reload() error {
	// TODO: This may lead to old monitors updating the guard before they're closed,
	// but this should not be an issue
	guard.stopAllMonitors()

	// TODO: Should we place a rollback mechanism here or is it best done elsewhere?
	err := guard.setupMonitors()
	if err != nil {
		return err
	}

	guard.startAllMonitors()

	return nil
}
