package core

type MonitorConfiguration interface {
	// Validate validates the configuration
	Validate() []error
}
