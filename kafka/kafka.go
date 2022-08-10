package kafka

type Kafka interface {
	NewProducer( /*...*/ )
	NewConsumer( /*...*/ )
}

/* TODO: WE ARE IMPLEMENTING THIS LIB */
type kafkaImpl struct {
}

func New() Kafka {
	return kafkaImpl{}
}
