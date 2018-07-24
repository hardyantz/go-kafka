package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	fmt.Println("sendingMessages")

	sendMessage()
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

func sendMessage() error {
	topic := "plankton"

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return err
	}

	go sendEvents(p)

	word := "sku " + strconv.Itoa(rand.Int())
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(word),
	}, nil)

	p.Flush(15 * 1000)

	return nil
}
