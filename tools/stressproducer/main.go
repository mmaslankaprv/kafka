package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_3_0_0
	client, err := sarama.NewClient([]string{"localhost:9091"}, config)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	err = client.RefreshMetadata()
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	brokers := client.Brokers()
	brokers[0].Open(config)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	req := sarama.MetadataRequest{
		AllowAutoTopicCreation: true,
		Topics:                 []string{"topic-1", "topic-2"},
		Version:                4,
	}

	resp, err := brokers[0].GetMetadata(&req)
	if err != nil {
		log.Printf("Error metadata response %v", err)
	}
	log.Printf("Metadata response %v", resp)
}
