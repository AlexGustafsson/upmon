package core

type WatchDelegate interface {
	// InterruptSignal is a channel which receives "true" when a monitor should stop watching
	InterruptSignal() <-chan bool
	// StatusUpdate is a channel which receives a service status
	StatusUpdate() chan<- ServiceStatus
}
