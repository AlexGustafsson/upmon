package core

import "sync"

type Check interface {
	// Name of the check
	Name() string
	// Description is the description of the check
	Description() string
	// Version of the monitor
	Version() string
	// Check the status of a service
	Perform() (Status, error)
	// Watch the status of a service continously
	Watch(update chan<- *ServiceStatus, stop <-chan bool, wg *sync.WaitGroup) error
	// Config is the configuration used for the monitor
	Config() MonitorConfiguration
}
