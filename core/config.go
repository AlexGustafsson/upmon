package core

// PeerConfig describes the configuration for a peer in the network
type PeerConfig struct {
  Certificate string
  Key string
  Fingerprint string
  Hostname string
  Port int
}

// ServiceConfig describes a service being monitored
type ServiceConfig struct {
  Hostname string
  Port int
  // Checks is an array of module names used for checking. The modules are run
  // in order. If one fails, the next is run.
  Checks []string
}

// Config describes a configuration used for Upmon
type Config struct {
  LogLevel string
  Peers map[string]PeerConfig
  Services map[string]ServiceConfig
}
