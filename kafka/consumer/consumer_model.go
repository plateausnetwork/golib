package consumer

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	Options struct {
		Configs            map[string]interface{}
		Topics             []string
		ReadMessageTimeout time.Duration
	}
	TopicPartition struct {
		Topic     string
		Partition int32
		Offset    int64
		Metadata  *string
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
		Timestamp      time.Time
	}
	Response struct {
		Error   error
		Message *Message
	}
)

func (o *Options) getConfigMap() (kafka.ConfigMap, error) {
	var cm = kafka.ConfigMap{}
	for key, val := range o.Configs {
		if err := cm.SetKey(key, val); err != nil {
			return nil, err
		}
	}
	return cm, nil
}
