package core

// MonitorConfiguration is a configuration for a monitor check
type MonitorConfiguration interface {
	// Validate validates the configuration
	Validate() []error
}
