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

type Message interface {
	Invalidates(other Message) bool
}

type Envelope struct {
	Sender      string
	MessageType MessageType
	Message     Message
}

type ConfigUpdateMessage struct {
	Version  int
	Services []configuration.ServiceConfiguration
}

type StatusUpdateMessage struct {
	Timestamp int64
	Node      string
	ServiceId string
	MonitorId string
	Status    monitoring.Status
}

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
		return message.Version > otherMessage.Version
	default:
		return false
	}
}

func (message *StatusUpdateMessage) Invalidates(other Message) bool {
	switch otherMessage := other.(type) {
	case *StatusUpdateMessage:
		return message.Timestamp > otherMessage.Timestamp && message.Node == otherMessage.Node && message.ServiceId == otherMessage.ServiceId && message.MonitorId == otherMessage.MonitorId
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
