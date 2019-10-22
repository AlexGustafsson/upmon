package core

// Peer describes the configuration for a peer in the network
type Peer struct {
  Certificate string
  Key string
  Fingerprint string
  Hostname string
  Port int
}

// Config describes a configuration used for Upmon
type Config struct {
  LogLevel string
  Peers map[string]Peer
}
