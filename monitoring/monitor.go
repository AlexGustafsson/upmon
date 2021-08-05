package monitoring

import (
	"sync"

	"github.com/AlexGustafsson/upmon/monitoring/core"
)

type Monitor struct {
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
	stop    chan bool
	wg      sync.WaitGroup
}

type Status core.Status

func (status Status) String() string {
	return core.Status(status).String()
}

type MonitorStatus struct {
	Err     error
	Status  Status
	Monitor *Monitor
}

func (monitor *Monitor) Watch(updates chan *MonitorStatus) error {
	upstreamUpdates := make(chan *core.ServiceStatus)
	stop := make(chan bool)
	monitor.stop = make(chan bool)

	err := monitor.Check.Watch(upstreamUpdates, stop, monitor.wg)
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
func (monitor *Monitor) Stop(wg sync.WaitGroup) {
	wg.Add(1)
	close(monitor.stop)
	monitor.wg.Wait()
	wg.Done()
}
