package ping

import (
	"os/exec"
	"time"

	"github.com/AlexGustafsson/upmon/monitor/core"
)

type Monitor struct {
}

func NewMonitor() (*Monitor, error) {
	return &Monitor{}, nil
}

func (monitor *Monitor) Name() string {
	return "ping"
}

func (monitor *Monitor) Description() string {
	return "Pinging monitor"
}

func (monitor *Monitor) Version() string {
	return "0.1.0"
}

func (monitor *Monitor) CheckImmediate(service core.Service) (core.ServiceStatus, error) {
	_, err := exec.Command("ping", service.Hostname(), "-c 1", "-t 1").Output()
	if err != nil {
		return core.StatusDown, nil
	}

	return core.StatusUp, nil
}

func (monitor *Monitor) Watch(service core.Service, interval time.Duration, delegate core.WatchDelegate) error {
	go func() {
		for {
			select {
			case <-delegate.InterruptSignal():
				return
			default:
				status, err := monitor.CheckImmediate(service)
				if err == nil {
					delegate.StatusUpdate() <- status
				}
			}
			time.Sleep(interval)
		}
	}()

	return nil
}
