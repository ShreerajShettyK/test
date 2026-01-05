package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

// KafkaConsumer represents a Kafka consumer
type KafkaConsumer struct {
	consumer sarama.ConsumerGroup
	topic    string
	groupID  string
}

// ConsumerGroupHandler handles messages consumed from Kafka
type ConsumerGroupHandler struct {
	messageCount int
	mu           sync.Mutex
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	log.Println("Consumer group session setup")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("Consumer group session cleanup")
	return nil
}

// ConsumeClaim processes messages from a topic partition
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE: Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine
	for message := range claim.Messages() {
		h.mu.Lock()
		h.messageCount++
		count := h.messageCount
		h.mu.Unlock()

		log.Printf("[Message #%d] Topic: %s, Partition: %d, Offset: %d, Key: %s, Value: %s, Timestamp: %v\n",
			count,
			message.Topic,
			message.Partition,
			message.Offset,
			string(message.Key),
			string(message.Value),
			message.Timestamp,
		)

		// Mark message as processed
		session.MarkMessage(message, "")
	}
	return nil
}

// NewKafkaConsumer creates a new Kafka consumer
func NewKafkaConsumer(brokers []string, groupID, topic string) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // Start from the oldest message

	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer group: %w", err)
	}

	return &KafkaConsumer{
		consumer: consumer,
		topic:    topic,
		groupID:  groupID,
	}, nil
}

// Start starts consuming messages
func (kc *KafkaConsumer) Start(ctx context.Context) error {
	handler := &ConsumerGroupHandler{}

	for {
		// `Consume` should be called inside an infinite loop
		// When a server-side rebalance happens, the consumer session will need to be recreated
		err := kc.consumer.Consume(ctx, []string{kc.topic}, handler)
		if err != nil {
			return fmt.Errorf("error from consumer: %w", err)
		}

		// Check if context was cancelled, signaling that the consumer should stop
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

// Close closes the consumer
func (kc *KafkaConsumer) Close() error {
	return kc.consumer.Close()
}

// RunConsumerDemo demonstrates consuming messages from Kafka
func RunConsumerDemo(brokers []string, groupID, topic string) {
	log.Println("=== Starting Kafka Consumer Demo ===")

	consumer, err := NewKafkaConsumer(brokers, groupID, topic)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	// Create context that listens for interrupt signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle SIGINT and SIGTERM
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	// Start consuming in a goroutine
	go func() {
		if err := consumer.Start(ctx); err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	}()

	log.Println("Consumer is running. Press Ctrl+C to stop...")

	// Wait for interrupt signal
	<-sigterm
	log.Println("\nReceived interrupt signal, shutting down...")
	cancel()

	log.Println("=== Consumer Demo Completed ===")
}
