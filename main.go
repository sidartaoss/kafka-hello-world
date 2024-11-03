package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// (1) Producer Configuration and Creation
	config := &kafka.ConfigMap{"bootstrap.servers": "host.docker.internal"}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	// (1) END

	// (2) Message Production
	topic := "test"
	msg := &kafka.Message{
		Value:          []byte("hi!"),
		Key:            []byte("1"),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	}

	wait := make(chan kafka.Event)
	err = producer.Produce(msg, wait)
	if err != nil {
		panic(err)
	}
	// (2) END

	<-wait
}
