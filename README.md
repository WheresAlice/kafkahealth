# kafkahealth

Connects to a Kafka topic and serves the timestamp of the latest message over http.

This isn't designed for real world usage, more a demonstration of the functionality of the [Benthos](https://github.com/Jeffail/benthos) stream processor.

## installation

`go install github.com/wheresalice/kafkahealth`

## usage

```
kafkahealth -b localhost:10092 -t test.topic [-p 8080]
curl localhost:8080/get
```

## http endpoints

* [/get](http://localhost:8080/get) - the latest timestamp
* [/get/stream](http://localhost:8080/get/stream) - a stream of timestamps
* [/get/ws](http://localhost:8080/get/ws) - a websocket of timestamps
