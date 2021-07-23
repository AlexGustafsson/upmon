package dns

import (
	"net"
	"time"

	"github.com/AlexGustafsson/upmon/monitor/core"
)

type Monitor struct {
	config  *MonitorConfiguration
	service core.Service
}

// NewMonitor creates a new monitor for a service
func NewMonitor(service core.Service, options map[string]interface{}) (*Monitor, error) {
	config, err := ParseConfiguration(options)
	if err != nil {
		return nil, err
	}

	return &Monitor{
		config:  config,
		service: service,
	}, nil
}

func (monitor *Monitor) Name() string {
	return "dns"
}

func (monitor *Monitor) Description() string {
	return "DNS resolver"
}

func (monitor *Monitor) Version() string {
	return "0.1.0"
}

func (monitor *Monitor) Service() core.Service {
	return monitor.service
}

func (monitor *Monitor) Config() core.MonitorConfiguration {
	return monitor.config
}

func (monitor *Monitor) CheckImmediate() (core.Status, error) {
	_, err := net.LookupIP(monitor.config.Hostname)
	if err != nil {
		return core.StatusUnknown, err
	}

	return core.StatusUp, nil
}

func (monitor *Monitor) Watch(update chan<- *core.ServiceStatus, stop <-chan bool) error {
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				status, err := monitor.CheckImmediate()
				update <- &core.ServiceStatus{
					Err:     err,
					Status:  status,
					Monitor: monitor,
				}
			}
			time.Sleep(monitor.config.Interval)
		}
	}()

	return nil
}
