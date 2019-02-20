package main

import pubsub "github.com/mellekoning/pubsubexample/pubsub"

func main() {
	print("hello world")
	pubsub := pubsub.New()
	pubsub.Publish("somekey", "value")
}
