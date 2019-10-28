package core

import (
  "io/ioutil"
  "github.com/BurntSushi/toml"
)

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

// GetUsedModules return a list of all used modules
func (config *Config) GetUsedModules() ([]string) {
  serviceMap := make(map[string]bool)
  for _, service := range config.Services {
    for _, module := range service.Checks {
      serviceMap[module] = true
    }
  }

  var services []string
  for service := range serviceMap {
    services = append(services, service)
  }

  return services
}

// LoadConfig reads and parses a given TOML configuration file
func LoadConfig(path string) (*Config, error) {
  LogDebug("Loading config from '%s'", path)
  dataBuffer, err := ioutil.ReadFile(path)
  if err != nil {
    return nil, err
  }

  data := string(dataBuffer)

  var config = new(Config)
  _, err = toml.Decode(data, config);
  if err != nil {
    return nil, err
  }

  return config, nil
}
