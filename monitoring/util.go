package monitoring

import (
	"fmt"

	"github.com/AlexGustafsson/upmon/monitoring/core"
	"github.com/AlexGustafsson/upmon/monitoring/dns"
	"github.com/AlexGustafsson/upmon/monitoring/ping"
)

// NewCheck creates a new check for a service by name
func NewCheck(checkName string, options map[string]interface{}) (core.Check, error) {
	switch checkName {
	case "ping":
		return ping.NewCheck(options)
	case "dns":
		return dns.NewCheck(options)
	default:
		return nil, fmt.Errorf("no such monitor")
	}
}
