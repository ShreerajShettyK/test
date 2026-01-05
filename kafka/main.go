package main

import (
	"flag"
	"log"
)

func main() {
	// Command line flags
	mode := flag.String("mode", "producer", "Mode: 'producer' or 'consumer'")
	brokers := flag.String("brokers", "localhost:9092", "Comma-separated list of Kafka brokers")
	topic := flag.String("topic", "test-topic", "Kafka topic name")
	groupID := flag.String("group", "test-consumer-group", "Consumer group ID (only for consumer mode)")
	flag.Parse()

	brokerList := []string{*brokers}

	log.Printf("Configuration:")
	log.Printf("  Mode: %s", *mode)
	log.Printf("  Brokers: %v", brokerList)
	log.Printf("  Topic: %s", *topic)
	if *mode == "consumer" {
		log.Printf("  Consumer Group: %s", *groupID)
	}
	log.Println()

	switch *mode {
	case "producer":
		RunProducerDemo(brokerList, *topic)
	case "consumer":
		RunConsumerDemo(brokerList, *groupID, *topic)
	default:
		log.Fatalf("Invalid mode: %s. Use 'producer' or 'consumer'", *mode)
	}
}
