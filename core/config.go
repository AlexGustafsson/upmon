package core

// PeerConfig describes the configuration for a peer in the network
type PeerConfig struct {
  Name string
  Description string
  Certificate string
  Key string
  Fingerprint string
  Hostname string
  Port int
}

// ServiceConfig describes a service being monitored
type ServiceConfig struct {
  Name string
  Description string
  Hostname string
  Port int
  // Checks is an array of module names used for checking. The modules are run
  // in order. If one fails, the next is run.
  Checks []string
}

// Config describes a configuration used for Upmon
type Config struct {
  LogLevel string
  Peers []PeerConfig
  Services []ServiceConfig
}

// GetPeerByName returns a peer if it is found, nil otherwise
func (config *Config) GetPeerByName(name string) (*PeerConfig) {
  for _, peer := range config.Peers {
    if peer.Name == name {
      return &peer
    }
  }

  return nil
}
