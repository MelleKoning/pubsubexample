# pubsub example



implementation of the pubsub example code from this website:
https://www.ardanlabs.com/blog/2016/11/application-focused-api-design.html

The goal here is to implement an interface as mock within a unit test. The mock implements the same methods as the underlying PubSub concrete struct.

With this mock, a certain error example of the pubsub 
can also be tested.

This example just shows that exposing a consumer of the pubsub-Struct can still be tested on those methods by exposing the underlying pubsub using its own constructor.

the constructor has to take in the pubsub dependency, this can be seen in the testcode as:

```
mockPublisher := &mockPublisher{}
c := consumer.NewConsumerInterface(mockPublisher)```




