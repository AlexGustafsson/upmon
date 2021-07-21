package core

import "time"

type Monitor interface {
	// Name of the monitor
	Name() string
	// Description of the monitor
	Description() string
	// Version of the monitor
	Version() string
	// Check the status of a service
	CheckImmediate(service Service) (ServiceStatus, error)
	// Watch the status of a service continously
	Watch(service Service, interval time.Duration, delegate WatchDelegate) error
}
