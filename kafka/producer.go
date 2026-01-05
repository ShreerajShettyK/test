package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// KafkaProducer demonstrates how to send messages to a Kafka topic
type KafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

// NewKafkaProducer creates a new Kafka producer
func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to acknowledge
	config.Producer.Retry.Max = 5                    // Retry up to 5 times
	config.Producer.Return.Successes = true          // Return success messages

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &KafkaProducer{
		producer: producer,
		topic:    topic,
	}, nil
}

// SendMessage sends a message to the Kafka topic
func (kp *KafkaProducer) SendMessage(key, value string) error {
	msg := &sarama.ProducerMessage{
		Topic: kp.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	partition, offset, err := kp.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	log.Printf("Message sent successfully! Topic: %s, Partition: %d, Offset: %d, Key: %s, Value: %s\n",
		kp.topic, partition, offset, key, value)
	return nil
}

// Close closes the producer
func (kp *KafkaProducer) Close() error {
	return kp.producer.Close()
}

// RunProducerDemo demonstrates sending messages to Kafka
func RunProducerDemo(brokers []string, topic string) {
	log.Println("=== Starting Kafka Producer Demo ===")

	producer, err := NewKafkaProducer(brokers, topic)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	// Send some sample messages
	messages := []struct {
		key   string
		value string
	}{
		{"user-1", "Hello from Kafka Producer!"},
		{"user-2", "This is message number 2"},
		{"user-3", "Learning Kafka topics"},
		{"user-1", "Another message for user-1"},
		{"user-4", "Final test message"},
	}

	for i, msg := range messages {
		err := producer.SendMessage(msg.key, msg.value)
		if err != nil {
			log.Printf("Error sending message %d: %v", i+1, err)
		}
		time.Sleep(1 * time.Second) // Wait a bit between messages
	}

	log.Println("=== Producer Demo Completed ===")
}
