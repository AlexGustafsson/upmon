package configuration

import (
	"fmt"

	"github.com/hashicorp/memberlist"
)

// PeerConfiguration configures a peer in the cluster
type PeerConfiguration struct {
	// Address is the address used for cluster communication
	Address string `koanf:"address"`
	// Port is the port used for cluster communication
	Port uint16 `koanf:"port"`
}

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
}

// Configuration of the service
type Configuration struct {
	// Name is the name of the node
	Name string `koanf:"name"`
	// Address is the address used for cluster communication
	Address string `koanf:"address"`
	// Port is the port used for cluster communication
	Port uint16 `koanf:"port"`
	// Peers is a list of peers that are part of the cluster
	Peers []PeerConfiguration `koanf:"peers"`
	// Services contains all the configured services, mapped by their name
	Services map[string]ServiceConfiguration `koanf:"services"`
	// Api contains configuration for the REST API
	Api struct {
		// Enabled controls whether or not the REST API is enabled
		Enabled bool `koanf:"enabled"`
		// Address is the address used to expose the API
		Address string `koanf:"address"`
		// Port is the port used to expose the API
		Port uint16 `koanf:"port"`
	} `koanf:"api"`
}

// MemberlistConfig creates a configuration for the memberlist gossip library
func (config *Configuration) MemberlistConfig() *memberlist.Config {
	memberlistConfig := memberlist.DefaultWANConfig()
	memberlistConfig.Name = config.Name
	memberlistConfig.BindAddr = config.Address
	memberlistConfig.BindPort = int(config.Port)
	return memberlistConfig
}

// PeerAddresses contains any configured peers' addresses
func (config *Configuration) PeerAddresses() []string {
	addresses := make([]string, len(config.Peers))
	for i, peer := range config.Peers {
		addresses[i] = fmt.Sprintf("%s:%d", peer.Address, peer.Port)
	}
	return addresses
}
