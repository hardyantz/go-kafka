package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strconv"
)

// set global variable for offset
var offset = 29

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          strconv.Itoa(offset),  // kafka will save group id, then if we change offset kafka still use this history offset
		"enable.auto.commit": "false",
	})

	if err != nil {
		panic(err)
	}

	topic := "plankton"

	c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	// set offset
	offset := kafka.Offset(offset)
	topicPartition := []kafka.TopicPartition{
		{Partition:0, Topic:&topic, Offset:offset},
	}

	// commit offset
	c.CommitOffsets(topicPartition)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
