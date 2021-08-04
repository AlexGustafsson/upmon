package clustering

//go:generate stringer -type=ClusterStatus
type ClusterStatus int

const (
	ClusterStatusHealthy ClusterStatus = iota
	ClusterStatusUnhealthy
)
