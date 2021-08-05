package clustering

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"

	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/monitoring"
	"github.com/hashicorp/memberlist"
	log "github.com/sirupsen/logrus"
)

type ClusterMember struct {
	ServicesVersion int
	Services        []configuration.ServiceConfiguration
}

type Cluster struct {
	self            string
	config          *configuration.Configuration
	Memberlist      *memberlist.Memberlist
	Members         map[string]*ClusterMember
	ServicesUpdates chan []configuration.ServiceConfiguration
}

type MessageType int

const (
	ConfigUpdate MessageType = iota
)

type Envelope struct {
	Sender      string
	MessageType MessageType
	Message     interface{}
}

type ConfigUpdateMessage struct {
	Version  int
	Services []configuration.ServiceConfiguration
}

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
		ServicesUpdates: make(chan []configuration.ServiceConfiguration),
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

	return cluster, nil
}

func (cluster *Cluster) Join(peers []string) error {
	_, err := cluster.Memberlist.Join(peers)
	return err
}

func (cluster *Cluster) welcome(node *memberlist.Node) {
	log.WithFields(log.Fields{"name": node.Name}).Info("welcoming node")

	// TODO: Only send public services
	envelope := &Envelope{
		Sender:      cluster.self,
		MessageType: ConfigUpdate,
		Message: ConfigUpdateMessage{
			Version:  0,
			Services: cluster.config.Services,
		},
	}

	data := bytes.NewBuffer(nil)
	gob.Register(ConfigUpdateMessage{})
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

func (cluster *Cluster) updateConfig(envelope *Envelope) {
	configUpdate := envelope.Message.(ConfigUpdateMessage)
	log.Debugf("Received config update message from '%s', version %d", envelope.Sender, configUpdate.Version)
	if member, ok := cluster.Members[envelope.Sender]; ok {
		if configUpdate.Version > member.ServicesVersion {
			member.Services = configUpdate.Services
			cluster.ServicesUpdates <- cluster.Services()
		}
	} else {
		cluster.Members[envelope.Sender] = &ClusterMember{
			ServicesVersion: configUpdate.Version,
			Services:        configUpdate.Services,
		}
		cluster.ServicesUpdates <- cluster.Services()
	}
}

// Services specifies the combined monitored services of the cluster
func (cluster *Cluster) Services() []configuration.ServiceConfiguration {
	services := make([]configuration.ServiceConfiguration, 0)

	for _, member := range cluster.Members {
		services = append(services, member.Services...)
	}

	return services
}

func (cluster *Cluster) BroadcastStatusUpdate(serviceId string, monitorId string, status monitoring.Status) error {
	log.WithFields(log.Fields{"service": serviceId, "monitor": monitorId, "status": status.String()}).Debugf("broadcasting status")
	return nil
}
