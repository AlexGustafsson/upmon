package configuration

import (
	"net"
	"strconv"

	"github.com/hashicorp/memberlist"
)

// MonitorConfiguration configures a monitor for a service
type MonitorConfiguration struct {
	// Type is the type of monitor, such as "ping" or "http"
	Type string `koanf:"type"`
	// Name is a name of the monitor
	Name string `koanf:"name"`
	// Description describes the monitor
	Description string `koanf:"description"`
	// Options contains arbitrary fields to use when configuring the monitor
	Options map[string]interface{} `koanf:"options"`
}

// ServiceConfiguration configures a service to be monitored
type ServiceConfiguration struct {
	// Description is a description of the service
	Description string `koanf:"description"`
	// Monitors is a list of monitors to use to monitor the service
	Monitors []MonitorConfiguration `koanf:"monitors"`
	// Private indicates whether or not the service configuration is shared with the cluster
	Private bool `koan:"private"`
}

// Configuration of the service
type Configuration struct {
	// Name is the name of the node
	Name string `koanf:"name"`
	// Bind is the address and port used for cluster communication
	Bind string `koanf:"bind"`
	// Peers is a list of peers' bind addresses and ports
	Peers []string `koanf:"peers"`
	// Services contains all the configured services, mapped by their name
	Services map[string]ServiceConfiguration `koanf:"services"`
	// Api contains configuration for the REST API
	Api struct {
		// Enabled controls whether or not the REST API is enabled
		Enabled bool `koanf:"enabled"`
		// Bind is the address and port used for cluster communication
		Bind string `koanf:"bind"`
	} `koanf:"api"`
}

// MemberlistConfig creates a configuration for the memberlist gossip library
func (config *Configuration) MemberlistConfig() (*memberlist.Config, error) {
	memberlistConfig := memberlist.DefaultWANConfig()
	memberlistConfig.Name = config.Name
	host, portString, err := net.SplitHostPort(config.Bind)
	if err != nil {
		return nil, err
	}
	port, err := strconv.ParseInt(portString, 10, 32)
	if err != nil {
		return nil, err
	}

	memberlistConfig.BindAddr = host
	memberlistConfig.BindPort = int(port)
	return memberlistConfig, nil
}
