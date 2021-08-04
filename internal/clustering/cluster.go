package clustering

import (
	"bytes"
	"encoding/gob"

	"github.com/hashicorp/memberlist"
	log "github.com/sirupsen/logrus"
)

type Cluster struct {
	Memberlist *memberlist.Memberlist
	self       string
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
	Version int
}

func NewCluster(config *memberlist.Config) (*Cluster, error) {
	cluster := &Cluster{
		self: config.Name,
	}

	delegate := &memberlistDelegate{
		cluster: cluster,
	}
	config.Events = delegate
	config.Conflict = delegate
	config.Delegate = delegate

	memberlist, err := memberlist.Create(config)
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

	envelope := &Envelope{
		Sender:      cluster.self,
		MessageType: ConfigUpdate,
		Message: ConfigUpdateMessage{
			Version: 0,
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
