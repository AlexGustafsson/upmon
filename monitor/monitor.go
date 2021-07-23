package monitor

import (
	"fmt"

	"github.com/AlexGustafsson/upmon/monitor/core"
	"github.com/AlexGustafsson/upmon/monitor/dns"
	"github.com/AlexGustafsson/upmon/monitor/ping"
)

// NewMonitor creates a new monitor for a service by name
func NewMonitor(name string, service core.Service, options map[string]interface{}) (core.Monitor, error) {
	switch name {
	case "ping":
		return ping.NewMonitor(service, options)
	case "dns":
		return dns.NewMonitor(service, options)
	default:
		return nil, fmt.Errorf("no such monitor")
	}
}
