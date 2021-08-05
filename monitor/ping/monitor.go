package ping

import (
	"sync"
	"time"

	"github.com/AlexGustafsson/upmon/monitor/core"
	"github.com/go-ping/ping"
	log "github.com/sirupsen/logrus"
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
	return "ping"
}

func (monitor *Monitor) Description() string {
	return "Pinging monitor"
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
	pinger := ping.New(monitor.config.Hostname)
	pinger.Count = monitor.config.Count
	pinger.Timeout = monitor.config.Timeout

	log.WithFields(log.Fields{"monitor": "ping"}).Debugf("resolving '%s'", monitor.config.Hostname)
	err := pinger.Resolve()
	if err != nil {
		return core.StatusUnknown, err
	}

	log.WithFields(log.Fields{"monitor": "ping"}).Debugf("pinging %s (%s)", pinger.IPAddr(), monitor.config.Hostname)
	err = pinger.Run()
	if err != nil {
		return core.StatusUnknown, err
	}

	statistics := pinger.Statistics()
	log.WithFields(log.Fields{"monitor": "ping"}).Debugf("sent %d packets, received %d packets, loss is %.0f%%", statistics.PacketsSent, statistics.PacketsRecv, statistics.PacketLoss)
	if statistics.PacketLoss == 0 {
		return core.StatusUp, nil
	}

	return core.StatusDown, nil
}

func (monitor *Monitor) Watch(update chan<- *core.ServiceStatus, stop <-chan bool, wg sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-stop:
				wg.Done()
				log.WithFields(log.Fields{"monitor": "ping"}).Debugf("exiting")
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
