package core

import "sync"

type Monitor interface {
	// Name of the monitor
	Name() string
	// Description of the monitor
	Description() string
	// Type of the monitor
	Type() string
	// TypeDescription is the description of the monitor type
	TypeDescription() string
	// Version of the monitor
	Version() string
	// Check the status of a service
	CheckImmediate() (Status, error)
	// Watch the status of a service continously
	Watch(update chan<- *ServiceStatus, stop <-chan bool, wg sync.WaitGroup) error
	// Service is the service the monitor monitors
	Service() Service
	// Config is the configuration used for the monitor
	Config() MonitorConfiguration
}

type DefaultMonitor struct {
	name        string
	description string
}

func (monitor *DefaultMonitor) Name() string {
	return monitor.name
}

func (monitor *DefaultMonitor) Description() string {
	return monitor.description
}
