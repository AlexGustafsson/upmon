package clustering

import (
	"bytes"
	"encoding/gob"

	"github.com/hashicorp/memberlist"
	log "github.com/sirupsen/logrus"
)

// memberlistDelegate is a delegate handling all of the events of the memberlist library
// see https://pkg.go.dev/github.com/hashicorp/memberlist
type memberlistDelegate struct {
	cluster *Cluster
}

func (delegate *memberlistDelegate) NotifyJoin(node *memberlist.Node) {
	// Ignore the self node
	if node.Name == delegate.cluster.self {
		return
	}

	log.WithFields(log.Fields{"name": node.Name, "address": node.Addr, "port": node.Port}).Info("node joined")
	if node.Name != delegate.cluster.self {
		delegate.cluster.welcome(node)
	}
}

func (delegate *memberlistDelegate) NotifyLeave(node *memberlist.Node) {
	log.WithFields(log.Fields{"name": node.Name, "address": node.Addr, "port": node.Port}).Info("node left")
}

func (delegate *memberlistDelegate) NotifyUpdate(node *memberlist.Node) {
	log.WithFields(log.Fields{"name": node.Name, "address": node.Addr, "port": node.Port}).Info("node updated")
}

func (delegate *memberlistDelegate) NotifyConflict(existing, other *memberlist.Node) {
	log.WithFields(log.Fields{"name": other.Name, "address": other.Addr, "port": other.Port}).Warning("node name conflict")
}

func (delegate *memberlistDelegate) NodeMeta(limit int) []byte {
	// No associated meta for now
	bytes := make([]byte, 0)
	return bytes
}

// RegisterGobNames must be called first
func (delegate *memberlistDelegate) NotifyMsg(data []byte) {
	reader := bytes.NewReader(data)
	encoder := gob.NewDecoder(reader)

	envelope := &Envelope{}
	err := encoder.Decode(envelope)
	if err != nil {
		log.Errorf("got unknown message: %v", err)
		return
	}

	switch envelope.MessageType {
	case ConfigUpdate:
		delegate.cluster.updateConfig(envelope.Message.(*ConfigUpdateMessage))
	case StatusUpdate:
		delegate.cluster.updateStatus(envelope.Message.(*StatusUpdateMessage))
	default:
		log.Errorf("got unknown message type from node '%s': %d", envelope.Sender, envelope.MessageType)
	}
}

func (delegate *memberlistDelegate) GetBroadcasts(overhead int, limit int) [][]byte {
	return delegate.cluster.broadcastQueue.GetBroadcasts(overhead, limit)
}

func (delegate *memberlistDelegate) LocalState(join bool) []byte {
	// No associated state for now
	bytes := make([]byte, 0)
	return bytes
}

func (delegate *memberlistDelegate) MergeRemoteState(buf []byte, join bool) {
	// No support for state merging for now
}
