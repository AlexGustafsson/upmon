package clustering

import (
	"bytes"
	"encoding/gob"

	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/hashicorp/memberlist"
	log "github.com/sirupsen/logrus"
)

type ClusterMember struct {
	ServicesVersion int
	Services        map[string]configuration.ServiceConfiguration
}

type Cluster struct {
	self       string
	config     *configuration.Configuration
	Memberlist *memberlist.Memberlist
	Members    map[string]*ClusterMember
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
	Services map[string]configuration.ServiceConfiguration
}

func NewCluster(config *configuration.Configuration) (*Cluster, error) {
	members := make(map[string]*ClusterMember)
	members[config.Name] = &ClusterMember{
		ServicesVersion: 0,
		Services:        config.Services,
	}

	cluster := &Cluster{
		self:    config.Name,
		config:  config,
		Members: members,
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

// Services specifies the combined monitored services of the cluster
func (cluster *Cluster) Services() map[string]configuration.ServiceConfiguration {
	services := make(map[string]configuration.ServiceConfiguration)

	for _, member := range cluster.Members {
		for name, service := range member.Services {
			if _, ok := services[name]; ok {
				// TODO: Merge
			} else {
				services[name] = service
			}
		}
	}

	return services
}
