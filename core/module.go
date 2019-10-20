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

// Host describes a network host or an application
type Host struct {
	Name        string
	Description string
	IP          string
	Port        int
}

type checkStatus func(*Host) (int, error)

// Module describes a module used for checking the status of a host
type Module struct {
	Name        string
	Description string
	Version     int
	CheckStatus checkStatus
}
