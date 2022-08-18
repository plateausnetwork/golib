package example

import (
	"github.com/rhizomplatform/golib/kafka/consumer"
	"github.com/rhizomplatform/golib/logger"
)

const TopicExample = "EXAMPLE"

func main() {
	c, err := consumer.New(consumer.Options{
		Topics: []string{TopicExample},
		Configs: map[string]interface{}{
			"bootstrap.servers": "localhost:9092",
			"auto.offset.reset": "earliest",
		},
	})
	if err != nil {
		logger.Error("failed to create consumer ", err)
	}
	for res := range c.Run() {
		if res.Error != nil {
			logger.Error("Response Message ERROR: ", res.Error)
			continue
		}
		logger.Info("MESSAGE: ", res.Message)
	}
	c.Close()
}
