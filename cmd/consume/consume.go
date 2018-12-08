package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
	"github.com/tkp-junnotantra/tc-nsq/messaging"
)

const (
	defaultConsumerMaxAttempts = 10
	defaultConsumerMaxInFlight = 100
)

func main() {
	// initiate consumer
	cfg := messaging.ConsumerConfig{
		Channel:       "", // TODO: update to given channel name
		LookupAddress: "", // TODO: update to nsqlookups adress
		Topic:         "", // TODO: update to given topic name
		MaxAttempts:   defaultConsumerMaxAttempts,
		MaxInFlight:   defaultConsumerMaxInFlight,
		Handler:       requeueMessage,
	}
	consumer := messaging.NewConsumer(cfg)

	// run consumer
	consumer.Run()

	// keep app alive until terminated
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-term:
		log.Println("Application terminated")
	}
}

func handleMessage(message *nsq.Message) error {
	// TODO: print and finish message
	return nil
}

func requeueMessage(message *nsq.Message) error {
	// TODO: requeue message
	return nil
}
