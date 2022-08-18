//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	Producer interface {
		Produce(msg Message) error
		Delivery() chan Delivery
		Flush(timeoutMs int) int
		Close()
	}
	producerImpl struct {
		kafka    *kafka.Producer
		delivery chan Delivery
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
	impl.deliveryHandler()
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

func (i producerImpl) Delivery() chan Delivery {
	return i.delivery
}

func (i producerImpl) Flush(timeoutMs int) int {
	return i.kafka.Flush(timeoutMs)
}

func (i producerImpl) Close() {
	close(i.delivery)
	i.kafka.Flush(0)
	i.kafka.Close()
}

func (i producerImpl) deliveryHandler() {
	i.delivery = make(chan Delivery)
	go func() {
		for e := range i.kafka.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				i.delivery <- Delivery{
					Error: ev.TopicPartition.Error,
					TopicPartition: TopicPartition{
						Topic:     *ev.TopicPartition.Topic,
						Partition: ev.TopicPartition.Partition,
						Offset:    int64(ev.TopicPartition.Offset),
					},
				}
			}
		}
	}()
}

/*
func Test() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "test"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
*/
