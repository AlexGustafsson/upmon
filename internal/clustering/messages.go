package clustering

import (
	"bytes"
	"encoding/gob"

	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/monitoring"
	"github.com/hashicorp/memberlist"
)

type MessageType int

const (
	ConfigUpdate MessageType = iota
	StatusUpdate
)

// Message is a message that may be sent across a cluster
type Message interface {
	// Invalidates specifies whether the message invalidates the other message
	Invalidates(other Message) bool
}

// Envelope contains a message to be sent across a cluster
type Envelope struct {
	// Sender is the node from which the message was sent
	Sender      string
	MessageType MessageType
	Message     Message
}

// ConfigUpdateMessage is sent from a node whenever there is a configuration update
type ConfigUpdateMessage struct {
	// Version is the incremental version of the configuration
	Version int
	// Node is the node for which this configuration is for
	Node     string
	Services []*configuration.ServiceConfiguration
}

// StatusUpdateMessage is broadcast from a node whenever there is a status change from a service
type StatusUpdateMessage struct {
	// Timestamp is the Unix timestamp of when the update occured
	Timestamp int64
	// Node is the node observing the status change
	Node string
	// Origin is the node which requested the monitoring
	Origin string
	// ServiceId is the id of the service which received an update
	ServiceId string
	// MonitorId is the id of the monitor which received an update
	MonitorId string
	// Status is the observed status
	Status monitoring.Status
}

// Broadcast is a message that may be broadcast across the cluster
type Broadcast struct {
	Envelope *Envelope
}

// RegisterGobNames must be called before
func (envelope *Envelope) Bytes() []byte {
	data := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(data)
	err := encoder.Encode(envelope)
	if err != nil {
		return nil
	}

	return data.Bytes()
}

func (message *ConfigUpdateMessage) Invalidates(other Message) bool {
	switch otherMessage := other.(type) {
	case *ConfigUpdateMessage:
		// Invalidate the other message if the version is higher than the other's (assuming same node)
		return message.Node == otherMessage.Node && message.Version > otherMessage.Version
	default:
		return false
	}
}

func (message *StatusUpdateMessage) Invalidates(other Message) bool {
	switch otherMessage := other.(type) {
	case *StatusUpdateMessage:
		// Invalidate the other message if the timestamp is newer than the other (assuming same monitor)
		return message.Timestamp > otherMessage.Timestamp && message.Origin == otherMessage.Origin && message.ServiceId == otherMessage.ServiceId && message.MonitorId == otherMessage.MonitorId
	default:
		return false
	}
}

func (broadcast *Broadcast) Invalidates(other memberlist.Broadcast) bool {
	otherBroadcast, ok := other.(*Broadcast)
	if !ok {
		return false
	}

	return broadcast.Envelope.Message.Invalidates(otherBroadcast.Envelope.Message)
}

func (broadcast *Broadcast) Message() []byte {
	return broadcast.Envelope.Bytes()
}

func (broadcast *Broadcast) Finished() {

}
