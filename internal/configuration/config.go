package configuration

import "github.com/hashicorp/memberlist"

type Peer struct {
	Address string `koanf:"address"`
	Port    uint16 `koanf:"port"`
}

type Configuration struct {
	Name    string `koanf:"name"`
	Address string `koanf:"address"`
	Port    uint16 `koanf:"port"`
	Peers   []Peer `koanf:"peers"`
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
