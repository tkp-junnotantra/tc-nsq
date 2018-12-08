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
		Channel:       "junno",          // TODO: update to desired value
		LookupAddress: "127.0.0.1:4161", // TODO: update to desired value
		Topic:         "top",            // TODO: update to desired value
		MaxAttempts:   defaultConsumerMaxAttempts,
		MaxInFlight:   defaultConsumerMaxInFlight,
		Handler:       handleMessage,
	}
	consumer := messaging.NewConsumer(cfg)

	cfg2 := messaging.ConsumerConfig{
		Channel:       "junno2",         // TODO: update to desired value
		LookupAddress: "127.0.0.1:4161", // TODO: update to desired value
		Topic:         "top",            // TODO: update to desired value
		MaxAttempts:   defaultConsumerMaxAttempts,
		MaxInFlight:   defaultConsumerMaxInFlight,
		Handler:       handleMessage2,
	}
	consumer2 := messaging.NewConsumer(cfg2)

	// run consumer
	consumer.Run()
	consumer2.Run()

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
	data := string(message.Body)
	log.Println("[1] consumed - " + data)
	message.Finish()
	return nil
}

func handleMessage2(message *nsq.Message) error {
	// TODO: print and finish message
	data := string(message.Body)
	log.Println("[2] consumed - " + data)
	message.Finish()
	return nil
}

func requeueMessage(message *nsq.Message) error {
	// TODO: requeue message
	return nil
}
