package consumer

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	Options struct {
		Configs map[string]interface{}
		Topics  []string
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
		Message Message
		Error   error
	}
)

func (o *Options) getConfigMap() (cm kafka.ConfigMap) {
	for key, val := range o.Configs {
		cm.SetKey(key, val)
	}
	return cm
}
