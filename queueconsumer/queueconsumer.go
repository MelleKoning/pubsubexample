package queueconsumer

import (
	pubsub "github.com/mellekoning/pubsubexample/pubsub"
)

// QueueConsumer implements the publisher
type QueueConsumer struct {
	Pub pubsub.PubSub
}

// NewConsumer return new instance of QueueConsumer, with a fresh
// publisher
func NewConsumer() *QueueConsumer {
	p := pubsub.New()
	return &QueueConsumer{
		Pub: *p,
	}
}

// Publisher interface shows the methods of
// the pubsub concrete struct
type Publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// QueueConsumer2 wraps the underlying
// publisher using an interface
// so that in the test, a mock of publisher
// can be used.
type QueueConsumer2 struct {
	Pub Publisher
}

// NewConsumerInterface to expose the pubsub as an
// own defined interface
func NewConsumerInterface(pub Publisher) *QueueConsumer2 {
	return &QueueConsumer2{
		Pub: pub,
	}
}
