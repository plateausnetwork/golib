package consumer

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	Consumer interface {
		Run(chResponse chan Response)
		Close() error
	}
	consumerImpl struct {
		kafka              *kafka.Consumer
		readMessageTimeout time.Duration
	}
)

func New(opts Options) (c Consumer, err error) {
	var impl consumerImpl
	config, err := opts.getConfigMap()
	if err != nil {
		return nil, err
	}
	if impl.kafka, err = kafka.NewConsumer(&config); err != nil {
		return nil, err
	}
	if err = impl.subscribeTopics(opts.Topics); err != nil {
		impl.Close()
		return nil, err
	}
	if opts.ReadMessageTimeout <= 0 {
		opts.ReadMessageTimeout = -1
	}
	impl.readMessageTimeout = opts.ReadMessageTimeout
	return impl, nil
}

func (i consumerImpl) Run(chResponse chan Response) {
	defer close(chResponse)
	for {
		i.kafka.ReadMessage(-1)
		msg, err := i.kafka.ReadMessage(i.readMessageTimeout)
		if err != nil {
			chResponse <- Response{
				Error:   err,
				Message: nil,
			}
			continue
		}
		if msg != nil {
			chResponse <- Response{
				Error:   nil,
				Message: i.convertMessage(msg),
			}
		}
	}
}

func (i consumerImpl) Close() error {
	return i.kafka.Close()
}

func (i consumerImpl) subscribeTopics(topics []string) error {
	return i.kafka.SubscribeTopics(topics, nil)
}

func (i consumerImpl) convertMessageHeaders(headers []kafka.Header) []Header {
	result := make([]Header, len(headers))
	for i, h := range headers {
		result[i] = Header{
			Key:   h.Key,
			Value: h.Value,
		}
	}
	return result
}

func (i consumerImpl) convertMessage(msg *kafka.Message) *Message {
	if msg == nil {
		return nil
	}
	return &Message{
		Key:       msg.Key,
		Value:     msg.Value,
		Headers:   i.convertMessageHeaders(msg.Headers),
		Timestamp: msg.Timestamp,
		TopicPartition: TopicPartition{
			Topic:     *msg.TopicPartition.Topic,
			Partition: msg.TopicPartition.Partition,
			Offset:    int64(msg.TopicPartition.Offset),
			Metadata:  msg.TopicPartition.Metadata,
		},
	}
}
