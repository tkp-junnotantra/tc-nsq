package messaging

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

const (
	defaultProducerPrefix = "tech_cur_"
)

type (
	ProducerConfig struct {
		NsqdAddress string
	}

	Producer struct {
		prod *nsq.Producer
	}
)

func NewProducer(cfg ProducerConfig) Producer {
	p, err := nsq.NewProducer(cfg.NsqdAddress, nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	return Producer{
		prod: p,
	}
}

func (p *Producer) Publish(topic string, msg interface{}) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	topic = defaultProducerPrefix + topic

	return p.prod.Publish(topic, payload)
}
