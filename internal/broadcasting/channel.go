package broadcasting

import "sync"

type SubscriptionCallback func()

type SignalChannel struct {
	sync.Mutex
	callbacks []SubscriptionCallback
	channels  []chan bool
}

func NewSignalChannel() *SignalChannel {
	return &SignalChannel{
		callbacks: make([]SubscriptionCallback, 0),
		channels:  make([]chan bool, 0),
	}
}

func (channel *SignalChannel) SubscribeCallback(callback SubscriptionCallback) {
	channel.Lock()
	defer channel.Unlock()

	channel.callbacks = append(channel.callbacks, callback)
}

func (channel *SignalChannel) SubscribeChannel(callback chan bool) {
	channel.Lock()
	defer channel.Unlock()

	channel.channels = append(channel.channels, callback)
}

func (channel *SignalChannel) Publish() {
	channel.Lock()
	defer channel.Unlock()

	for _, callback := range channel.callbacks {
		callback()
	}

	for _, callback := range channel.channels {
		callback <- true
	}
}
