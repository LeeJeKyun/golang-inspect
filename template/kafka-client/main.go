package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

type Config struct {
	BootstrapServers string `json:"bootstrap.servers"`
	Topic            string `json:"topic"`
}

func main() {

	var config Config

	//설정파일 open
	file, err := os.Open("./config.json")
	if err != nil {
		fmt.Printf("Error opening config file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	//config 구조체 생성
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)

	//Producer 생성
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}
	defer p.Close()

	//Consumer 생성
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
		"group.id":          "test",
	})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
	}
	defer c.Close()

	//구독 토픽 설정
	err = c.SubscribeTopics([]string{config.Topic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topic: %s\n", err)
		os.Exit(1)
	}

	//Consumer 고루틴 생성
	go func() {
		for {
			msg, err := c.ReadMessage(-1)
			if err != nil {
				fmt.Printf("Error reading message: %v\n", err)
				continue
			}
			fmt.Printf("Consumer event for topic %s: key = %-10s value = %s\n", *msg.TopicPartition.Topic, string(msg.Key), string(msg.Value))
		}
	}()

	topic := config.Topic
	factor := true
	for factor {
		var input string
		fmt.Scanln(&input)
		if input == "exit" {
			factor = false
		} else {
			go p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(input),
			}, nil)
		}
	}

}
