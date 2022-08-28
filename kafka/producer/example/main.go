package example

import (
	"encoding/json"

	"github.com/rhizomplatform/golib/kafka/producer"
	"github.com/rhizomplatform/golib/logger"
)

const TopicExample = "EXAMPLE"

func main() {
	p, err := producer.New(producer.Options{
		Configs: map[string]interface{}{
			"bootstrap.servers": "localhost:9092",
			"acks":              "all",
		},
	})
	if err != nil {
		logger.Fatal("failed to create producer ", err)
	}
	msg, err := makeExampleMessage()
	if err != nil {
		logger.Fatal("failed to generate msg ", err)
	}
	if err = p.Produce(*msg); err != nil {
		logger.Fatal("failed to produce msg ", err)
	}
	for delivery := range p.Delivery() {
		if delivery.Error != nil {
			logger.Error("DELIVERY ERROR: ", delivery.Error)
			continue
		}
		logger.Info("DELIVERY: ", delivery.TopicPartition)
	}
	p.Close()
}

func makeExampleMessage() (*producer.Message, error) {
	msgBytes, err := json.Marshal(map[string]string{
		"name":  "rhz",
		"email": "contact@rhizom.me",
	})
	if err != nil {
		logger.Error("failed to convert json ", err)
		return nil, err
	}
	msgHeader := []producer.Header{
		{Key: "Content-Type", Value: []byte("application/json")},
	}
	msg := producer.Message{
		TopicPartition: producer.TopicPartition{
			Topic: TopicExample,
		},
		Value:   msgBytes,
		Headers: msgHeader,
	}
	return &msg, nil
}
