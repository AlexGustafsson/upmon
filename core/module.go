package core

// The status type of a host
const (
	StatusUnknown   = iota
	StatusDown      = iota
	StatusGoingUp   = iota
	StatusUp        = iota
	StatusGoingDown = iota
)

// Module version
const (
	ModuleVersionUnknown = iota
	ModuleVersion1       = iota
)

// ServiceInfo describes information gathered when checking the status of a service
type ServiceInfo struct {
	Status int
}

type checkService func(*ServiceConfig) (*ServiceInfo, error)

// Module describes a module used for checking the status of a host
type Module struct {
	Name         string
	Description  string
	Version      int
	CheckService checkService
}
