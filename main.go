package main

import (
	"github.com/Jeffail/benthos/lib/stream"
	"github.com/Jeffail/benthos/lib/input"
	"github.com/Jeffail/benthos/lib/types"
	"time"
	"os"
	"os/signal"
	"syscall"
	"log"
	"github.com/Jeffail/benthos/lib/output"
	"github.com/jessevdk/go-flags"
	"fmt"
	"bytes"
)

var opts struct {
	Broker string `short:"b" long:"broker" description:"kafka broker to connect to" required:"true"`
	Topic string `short:"t" long:"topic" description:"kafka topic to check" required:"true"`
	Port int `short:"p" long:"port" description:"http port to listen on" default:"8080"`
}
type TimeStamp struct{}

func (TimeStamp) ProcessMessage(m types.Message) ([]types.Message, types.Response) {
	// Create a copy of the original message
	result := m.Copy()

	// For each message part replace its contents with the timestamp we received it at
	result.Iter(func(i int, part types.Part) error {
		var buf bytes.Buffer

		buf.WriteString(m.CreatedAt().String())

		part.Set(buf.Bytes())
		return nil
	})

	return []types.Message{result}, nil
}



func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	conf := stream.NewConfig()

	conf.Input.Type = input.TypeKafka
	conf.Input.Kafka.Addresses = []string{opts.Broker}
	conf.Input.Kafka.Topic = opts.Topic

	conf.Output.Type = output.TypeHTTPServer
	address := fmt.Sprintf("0.0.0.0:%d", opts.Port)
	conf.Output.HTTPServer.Address = address

	log.Printf("connecting to: %v", conf.Input.Kafka.Addresses)
	log.Printf("reading from topic: %v", conf.Input.Kafka.Topic)
	log.Printf("listening on: %s", address)

	s, err := stream.New(conf, stream.OptAddProcessors(func() (types.Processor, error) {
		return TimeStamp{}, nil
	}))
	if err != nil {
		panic(err)
	}

	defer s.Stop(time.Second)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
		log.Println("Received SIGTERM, the service is closing")
	}



}