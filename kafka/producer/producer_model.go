package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	Options struct {
		Configs map[string]interface{}
	}
	TopicPartition struct {
		Topic     string
		Partition int32
		Offset    int64
	}
	Header struct {
		Key   string
		Value []byte
	}
	Message struct {
		TopicPartition TopicPartition
		Value          []byte
		Key            []byte
		Headers        []Header
	}
	Delivery struct {
		TopicPartition TopicPartition
		Error          error
	}
)

func (o *Options) getConfigMap() (cm kafka.ConfigMap) {
	for key, val := range o.Configs {
		cm.SetKey(key, val)
	}
	return cm
}

func (m Message) kafkaHeaders() []kafka.Header {
	headers := make([]kafka.Header, len(m.Headers))
	for i, h := range m.Headers {
		headers[i] = kafka.Header{
			Key:   h.Key,
			Value: h.Value,
		}
	}
	return headers
}

func (m Message) kafkaTopicPartition() (tp kafka.TopicPartition) {
	tp.Topic = &m.TopicPartition.Topic
	tp.Partition = m.TopicPartition.Partition
	tp.Offset.Set(m.TopicPartition.Offset)
	return tp
}
