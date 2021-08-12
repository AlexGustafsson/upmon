package clustering

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"time"

	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/internal/storage"
	"github.com/AlexGustafsson/upmon/monitoring"
	"github.com/hashicorp/memberlist"
	log "github.com/sirupsen/logrus"
)

// ClusterMember is a member of a cluster
type ClusterMember struct {
	// ServicesVersion is the current (incremental) version of the configured services
	ServicesVersion int
	// Services are the services configured for the cluster member
	Services []*configuration.ServiceConfiguration
}

// Cluster is a cluster of monitoring nodes
type Cluster struct {
	// self is the name / id of the current node
	self   string
	config *configuration.Configuration
	// broadcastQueue is the queue used for broadcasting messages across the cluster
	broadcastQueue memberlist.TransmitLimitedQueue
	Memberlist     *memberlist.Memberlist
	// Members are all welcomed members of the cluster from this node's point of view
	Members map[string]*ClusterMember
	// ServiceUpdates published the configuration any time it's updated
	ServicesUpdates chan []*configuration.ServiceConfiguration
	// Store is a state store for the cluster
	Store *storage.Store
}

// NewCluster creates a new cluster based on the given configuration
func NewCluster(config *configuration.Configuration) (*Cluster, error) {
	members := make(map[string]*ClusterMember)
	members[config.Name] = &ClusterMember{
		ServicesVersion: 0,
		Services:        config.Services,
	}

	cluster := &Cluster{
		self:            config.Name,
		config:          config,
		Members:         members,
		ServicesUpdates: make(chan []*configuration.ServiceConfiguration),
		Store:           storage.NewStore(),
	}

	memberlistConfig, err := config.MemberlistConfig()
	if err != nil {
		return nil, err
	}

	delegate := &memberlistDelegate{
		cluster: cluster,
	}
	memberlistConfig.Events = delegate
	memberlistConfig.Conflict = delegate
	memberlistConfig.Delegate = delegate
	// Disable memberlist logging
	memberlistConfig.LogOutput = ioutil.Discard

	memberlist, err := memberlist.Create(memberlistConfig)
	if err != nil {
		return nil, err
	}

	cluster.Memberlist = memberlist

	cluster.broadcastQueue.NumNodes = cluster.Memberlist.NumMembers
	cluster.broadcastQueue.RetransmitMult = 1

	return cluster, nil
}

// Join joins one or more peers using the given addresses
func (cluster *Cluster) Join(peers []string) error {
	_, err := cluster.Memberlist.Join(peers)
	return err
}

// welcome welcomes another node by publishing the current node's configuration
func (cluster *Cluster) welcome(node *memberlist.Node) {
	log.WithFields(log.Fields{"name": node.Name}).Info("welcoming node")

	// Only distribute public services
	publicServices := make([]*configuration.ServiceConfiguration, 0)
	for _, service := range cluster.config.Services {
		if !service.Private {
			publicServices = append(publicServices, service)
		}
	}

	envelope := &Envelope{
		Sender:      cluster.self,
		MessageType: ConfigUpdate,
		Message: &ConfigUpdateMessage{
			Node:     cluster.self,
			Version:  0,
			Services: publicServices,
		},
	}

	data := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(data)
	err := encoder.Encode(envelope)
	if err != nil {
		log.WithFields(log.Fields{"name": node.Name}).Errorf("failed to welcome node: %v", err)
		return
	}

	err = cluster.Memberlist.SendReliable(node, data.Bytes())
	if err != nil {
		log.WithFields(log.Fields{"name": node.Name}).Errorf("failed to welcome node: %v", err)
		return
	}
}

func (cluster *Cluster) Status() ClusterStatus {
	health := cluster.Memberlist.GetHealthScore()
	if health == 0 {
		return ClusterStatusHealthy
	}

	return ClusterStatusUnhealthy
}

// updateConfig handles incoming configuration updates from the cluster
func (cluster *Cluster) updateConfig(configUpdate *ConfigUpdateMessage) {
	log.Debugf("Received config update message from '%s', version %d", configUpdate.Node, configUpdate.Version)
	if member, ok := cluster.Members[configUpdate.Node]; ok {
		if configUpdate.Version > member.ServicesVersion {
			member.Services = configUpdate.Services
			cluster.ServicesUpdates <- cluster.Services()
		}
	} else {
		cluster.Members[configUpdate.Node] = &ClusterMember{
			ServicesVersion: configUpdate.Version,
			Services:        configUpdate.Services,
		}
		cluster.ServicesUpdates <- cluster.Services()
	}
}

// Services specifies the combined monitored services of the cluster
func (cluster *Cluster) Services() []*configuration.ServiceConfiguration {
	services := make([]*configuration.ServiceConfiguration, 0)

	for _, member := range cluster.Members {
		services = append(services, member.Services...)
	}

	return services
}

// BroadcastStatusUpdate broadcasts an update for a specified monitor over an eventual consistent channel
func (cluster *Cluster) BroadcastStatusUpdate(origin string, serviceId string, monitorId string, status monitoring.Status) error {
	log.WithFields(log.Fields{"service": serviceId, "monitor": monitorId, "status": status.String()}).Debugf("broadcasting status")
	envelope := &Envelope{
		Sender:      cluster.self,
		MessageType: StatusUpdate,
		Message: &StatusUpdateMessage{
			Timestamp: time.Now().Unix(),
			Node:      cluster.self,
			Origin:    origin,
			ServiceId: serviceId,
			MonitorId: monitorId,
			Status:    status,
		},
	}

	data := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(data)
	err := encoder.Encode(envelope)
	if err != nil {
		return err
	}

	broadcast := &Broadcast{
		Envelope: envelope,
	}

	cluster.broadcastQueue.QueueBroadcast(broadcast)

	return nil
}

// updateStatus handles incoming status updates broadcasted on the cluster
func (cluster *Cluster) updateStatus(statusUpdate *StatusUpdateMessage) {
	log.Debugf("Received status update message from '%s' for node '%s', service %s, monitor %s: %s", statusUpdate.Node, statusUpdate.Origin, statusUpdate.ServiceId, statusUpdate.MonitorId, statusUpdate.Status.String())
	cluster.Store.AssertOrigin(statusUpdate.Origin).AssertService(statusUpdate.ServiceId).AssertMonitor(statusUpdate.MonitorId).SetStatusForNode(statusUpdate.Node, statusUpdate.Status)
}
