package queueconsumer_test

import (
	"errors"
	"testing"

	consumer "github.com/mellekoning/pubsubexample/queueconsumer"
)

func TestQueueConsumerCreationSuccess(t *testing.T) {

	// This test just runs against
	// the actual PubSub struct,
	// so no interface is used

	// Arrange
	c := consumer.NewConsumer()

	// Act
	c.Pub.Publish("key", "value")

}

type mockPublisher struct {
}

// Publish sends the data for the specified key.
func (ps *mockPublisher) Publish(key string, v interface{}) error {
	/* impl */

	if key == "error" {
		return errors.New("some error occurred")
	}

	println("publish")
	return nil
}

// Subscribe requests the data for the specified key.
func (ps *mockPublisher) Subscribe(key string) error {
	/* impl */
	println("subscribe")
	return nil
}

func TestQueueConsumer2CreationSuccess(t *testing.T) {

	// This test uses a mock of the underlying
	// Publisher. some expectations could
	// be setup in the test

	// Arrange
	mockPublisher := &mockPublisher{}
	c := consumer.NewConsumerInterface(mockPublisher)

	// Act
	err := c.Pub.Publish("error", "value")

	if err == nil {
		t.Fail() // wee assumed an error
	}
}
