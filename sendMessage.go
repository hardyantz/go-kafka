package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hardyantz/go-kafka/libraries"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("usage ./sendMessage <topic> <message>")
		os.Exit(1)
	}

	topic := os.Args[1]
	message := os.Args[2]

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("create producer %v \n", p)

	go libraries.SendEvents(p)

	err = libraries.SendMessage(message, topic)
	if err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(0)
}
