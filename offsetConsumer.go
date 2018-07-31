package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	var topic = "plankton"  // set topic
	var offsetNumber = 29 // set offset number
	var groupId = "plankton-group1" //  set group id

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost",
		"group.id":           groupId, // kafka will save group id, then if we change offset kafka still use this history offset
		"enable.auto.commit": "false",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	// set offset
	offset := kafka.Offset(offsetNumber)
	topicPartition := []kafka.TopicPartition{
		{Partition: 0, Topic: &topic, Offset: offset},
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
