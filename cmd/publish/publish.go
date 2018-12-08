package main

import (
	"github.com/tkp-junnotantra/tc-nsq/messaging"
)

func main() {
	// initiate producer
	prodConf := messaging.ProducerConfig{
		NsqdAddress: "", // TODO: update to nsqd address
	}
	prod := messaging.NewProducer(prodConf)

	// publish message
	topic := "" // TODO: update to given topic name
	msg := ""   // TODO: write your message here
	prod.Publish(topic, msg)
}
