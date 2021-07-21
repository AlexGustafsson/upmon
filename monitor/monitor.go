package monitor

import (
	"fmt"

	"github.com/AlexGustafsson/upmon/monitor/irc"
	"github.com/AlexGustafsson/upmon/monitor/ping"
)

func NewMonitor(name string) (monitor.Monitor, error) {
	switch name {
	case "irc":
		return irc.NewMonitor()
	case "ping":
		return ping.NewMonitor()
	default:
		return nil, fmt.Errorf("no such monitor")
	}
}
