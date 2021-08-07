package monitoring

import (
	"sync"

	"github.com/AlexGustafsson/upmon/monitoring/core"
)

// Monitor is a service monitor
type Monitor struct {
	stop chan bool
	wg   sync.WaitGroup
	// Id is an identifier of the monitor, unique for the service
	Id string
	// Name of the monitor
	Name string
	// Description of the monitor
	Description string
	// Check is the check performed by the monitor
	Check core.Check
	// Service is the service the check is performed on
	Service *Service
}

type Status core.Status

const (
	StatusUp Status = iota
	StatusTransitioningUp
	StatusTransitioningDown
	StatusDown
	StatusUnknown
)

func (status Status) String() string {
	return core.Status(status).String()
}

// MonitorStatus is a status update from a monitor
type MonitorStatus struct {
	Err     error
	Status  Status
	Monitor *Monitor
}

// Watch watches a monitor for changes which are published on the given channel
func (monitor *Monitor) Watch(updates chan *MonitorStatus) error {
	upstreamUpdates := make(chan *core.ServiceStatus)
	stop := make(chan bool)
	monitor.stop = make(chan bool)

	err := monitor.Check.Watch(upstreamUpdates, stop, &monitor.wg)
	if err != nil {
		return err
	}

	monitor.wg.Add(1)
	go func() {
		for update := range upstreamUpdates {
			updates <- &MonitorStatus{
				Err:     update.Err,
				Status:  Status(update.Status),
				Monitor: monitor,
			}
		}
		monitor.wg.Done()
	}()

	return nil
}

// Stop stops the monitor, blocking until it's stopped
func (monitor *Monitor) Stop(wg *sync.WaitGroup) {
	wg.Add(1)
	close(monitor.stop)
	monitor.wg.Wait()
	wg.Done()
}
