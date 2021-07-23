package configuration

import "github.com/hashicorp/memberlist"

type PeerConfiguration struct {
	Address string `koanf:"address"`
	Port    uint16 `koanf:"port"`
}

type CheckConfiguration struct {
	Type        string                 `koanf:"type"`
	Name        string                 `koanf:"name"`
	Description string                 `koanf:"description"`
	Options     map[string]interface{} `koanf:"options"`
}

type ServiceConfiguration struct {
	Description string               `koanf:"description"`
	Checks      []CheckConfiguration `koanf:"checks"`
}

type Configuration struct {
	Name     string                          `koanf:"name"`
	Address  string                          `koanf:"address"`
	Port     uint16                          `koanf:"port"`
	Peers    []PeerConfiguration             `koanf:"peers"`
	Services map[string]ServiceConfiguration `koanf:"services"`
	Api      struct {
		Enabled bool   `koanf:"enabled"`
		Address string `koanf:"address"`
		Port    uint16 `koanf:"port"`
	} `koanf:"api"`
}

func (config *Configuration) MemberlistConfig() *memberlist.Config {
	memberlistConfig := memberlist.DefaultWANConfig()
	memberlistConfig.Name = config.Name
	memberlistConfig.BindAddr = config.Address
	memberlistConfig.BindPort = int(config.Port)
	return memberlistConfig
}

func (config *Configuration) PeerAddresses() []string {
	addresses := make([]string, len(config.Peers))
	for i, peer := range config.Peers {
		addresses[i] = peer.Address
	}
	return addresses
}
