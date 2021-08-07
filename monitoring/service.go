package monitoring

// Service is a monitored service
type Service struct {
	// Id is an identifier of the monitor, unique for all services
	Id string
	// Name of the monitor
	Name string
	// Description of the monitor
	Description string
	// Origin is the node from which this service is configured
	Origin string
}
