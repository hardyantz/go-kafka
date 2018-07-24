package main

import (
	"fmt"
			"github.com/confluentinc/confluent-kafka-go/kafka"
)

var topic = "plankton"

func main() {
	fmt.Println("sendingMessages")

	message := "test send message"

	err := sendMessage(message)
	if err != nil {
		fmt.Println(err)
	}

}

func sendEvents(p *kafka.Producer) {
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
}

func sendMessage(message string) error {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return err
	}

	go sendEvents(p)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		return err
	}

	p.Flush(15 * 1000)

	return nil
}
