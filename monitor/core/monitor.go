package core

type Monitor interface {
	// Name of the monitor
	Name() string
	// Description of the monitor
	Description() string
	// Version of the monitor
	Version() string
	// Check the status of a service
	CheckImmediate() (Status, error)
	// Watch the status of a service continously
	Watch(update chan<- *ServiceStatus, stop <-chan bool) error
	// Service is the service the monitor monitors
	Service() Service
	// Config is the configuration used for the monitor
	Config() MonitorConfiguration
}
