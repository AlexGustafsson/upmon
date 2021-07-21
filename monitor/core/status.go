package core

type ServiceStatus int

const (
	StatusUp ServiceStatus = iota
	StatusTransitioningUp
	StatusTransitioningDown
	StatusDown
	StatusUnknown
)
