package core

//go:generate stringer -type=Status
type Status int

const (
	StatusUp Status = iota
	StatusTransitioningUp
	StatusTransitioningDown
	StatusDown
	StatusUnknown
)

type ServiceStatus struct {
	Err    error
	Status Status
	Check  Check
}
