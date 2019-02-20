package pubsub

// Package pubsub simulates a package that provides
// publication/subscription type services.

// PubSub provides access to a queue system.
type PubSub struct {
	/* impl */
}

// New creates a pubsub value for use.
func New( /* impl */ ) *PubSub {
	ps := PubSub{
		/* impl */
	}

	/* impl */
	return &ps
}

// Publish sends the data for the specified key.
func (ps *PubSub) Publish(key string, v interface{}) error {
	/* impl */
	println("publish")
	return nil
}

// Subscribe requests the data for the specified key.
func (ps *PubSub) Subscribe(key string) error {
	/* impl */
	println("subscribe")
	return nil
}
