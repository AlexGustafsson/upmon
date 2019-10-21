package core

// Peer describes the configuration for a peer in the network
type Peer struct {
  PrivateKey string
  PublicKey string
  Hostname string
  Port int
}

// Config describes a configuration used for Upmon
type Config struct {
  LogLevel string
  Peers map[string]Peer
}
