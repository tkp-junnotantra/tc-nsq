package main

import (
	"github.com/tkp-junnotantra/tc-nsq/messaging"
)

func main() {
	// initiate producer
	prodConf := messaging.ProducerConfig{
		NsqdAddress: "127.0.0.1:4150",
	}
	prod := messaging.NewProducer(prodConf)

	// publish message
	topic := "top"
	msg := "hello world" // TODO: write your message here
	prod.Publish(topic, msg)
}
