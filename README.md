# kafkahealth

Connects to a Kafka topic and serves the timestamp of the latest message over http.

## usage

```
kafkahealth -b localhost:10092 -t test.topic [-p 8080]
curl localhost:8080/get
```

## endpoints

* [/get](http://localhost:8080/get) - the latest timestamp
* [/get/stream](http://localhost:8080/get/stream) - a stream of timestamps
* [/get/ws](http://localhost:8080/get/ws) - a websocket of timestamps
