//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	Producer interface {
		Produce(msg Message) error
		Delivery(chDelivery chan Delivery)
		Flush(timeoutMs int) int
		Close()
	}
	producerImpl struct {
		kafka *kafka.Producer
	}
)

func New(opts Options) (p Producer, err error) {
	var (
		config = opts.getConfigMap()
		impl   = producerImpl{}
	)
	if impl.kafka, err = kafka.NewProducer(&config); err != nil {
		return nil, err
	}
	return impl, nil
}

func (i producerImpl) Produce(msg Message) error {
	return i.kafka.Produce(
		&kafka.Message{
			TopicPartition: msg.kafkaTopicPartition(),
			Value:          msg.Value,
			Key:            msg.Key,
			Headers:        msg.kafkaHeaders(),
		}, nil)
}

func (i producerImpl) Delivery(chDelivery chan Delivery) {
	defer close(chDelivery)
	for e := range i.kafka.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			chDelivery <- Delivery{
				Error: ev.TopicPartition.Error,
				TopicPartition: TopicPartition{
					Topic:     *ev.TopicPartition.Topic,
					Partition: ev.TopicPartition.Partition,
					Offset:    int64(ev.TopicPartition.Offset),
				},
			}
		}
	}
}

func (i producerImpl) Flush(timeoutMs int) int {
	return i.kafka.Flush(timeoutMs)
}

func (i producerImpl) Close() {
	i.kafka.Flush(0)
	i.kafka.Close()
}
