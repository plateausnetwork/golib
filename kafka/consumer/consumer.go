package consumer

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	Consumer interface {
		Run() chan Response
		Close() error
	}
	consumerImpl struct {
		kafka              *kafka.Consumer
		messages           chan Response
		readMessageTimeout time.Duration
	}
)

func New(opts Options) (c Consumer, err error) {
	var (
		config = opts.getConfigMap()
		impl   = consumerImpl{}
	)
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

func (i consumerImpl) Run() chan Response {
	i.messages = make(chan Response)
	for {
		msg, err := i.kafka.ReadMessage(i.readMessageTimeout)
		i.messages <- Response{
			Error: err,
			Message: Message{
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
			},
		}
	}
}

func (i consumerImpl) Close() error {
	close(i.messages)
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
