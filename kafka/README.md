# Kafka Producer and Consumer Example

This is a basic implementation of Kafka producer and consumer in Go to help understand Kafka topics.

## Overview

### What is a Kafka Topic?

A **Kafka topic** is a category or feed name to which records are published. Topics in Kafka are always multi-subscriber; that is, a topic can have zero, one, or many consumers that subscribe to the data written to it.

Key concepts:
- **Topic**: A stream of messages of a particular type
- **Partition**: Topics are split into partitions for scalability
- **Offset**: Each message within a partition has a unique offset
- **Producer**: Publishes messages to topics
- **Consumer**: Subscribes to topics and processes messages
- **Consumer Group**: Multiple consumers working together to consume from a topic

## Architecture

```
Producer → [Kafka Topic with Partitions] → Consumer Group
              ├── Partition 0 (messages with offsets)
              ├── Partition 1 (messages with offsets)
              └── Partition 2 (messages with offsets)
```

## Files

- `producer.go` - Kafka producer implementation that sends messages to a topic
- `consumer.go` - Kafka consumer implementation that reads messages from a topic
- `main.go` - Main entry point with command-line options

## Prerequisites

1. Docker installed
2. Go installed (1.21 or later)

## Setup Kafka with Docker

### 1. Create Docker Network (if not exists)
```powershell
docker network create kafka-network
```

### 2. Start Zookeeper (Kafka dependency)
```powershell
docker run -d --name zookeeper --network kafka-network -p 2181:2181 -e ZOOKEEPER_CLIENT_PORT=2181 confluentinc/cp-zookeeper:latest
```

### 3. Start Kafka
```powershell
docker run -d --name kafka --network kafka-network -p 9092:9092 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 confluentinc/cp-kafka:latest
```

### 4. Verify containers are running
```powershell
docker ps
```

## Usage

### Install Dependencies
```powershell
go mod download
```

### Run Producer
Sends 5 sample messages to the Kafka topic:
```powershell
go run . -mode producer -topic test-topic -brokers localhost:9092
```

### Run Consumer
Consumes messages from the Kafka topic:
```powershell
go run . -mode consumer -topic test-topic -group test-consumer-group -brokers localhost:9092
```

### Running Both (in separate terminals)

**Terminal 1 - Start Consumer:**
```powershell
go run . -mode consumer
```

**Terminal 2 - Run Producer:**
```powershell
go run . -mode producer
```

## Command Line Options

- `-mode`: `producer` or `consumer` (default: `producer`)
- `-brokers`: Kafka broker addresses (default: `localhost:9092`)
- `-topic`: Topic name (default: `test-topic`)
- `-group`: Consumer group ID (default: `test-consumer-group`, only for consumer)

## How It Works

### Producer
1. Connects to Kafka broker
2. Creates messages with keys and values
3. Sends messages to the specified topic
4. Kafka assigns messages to partitions (based on key hash)
5. Returns partition and offset information

### Consumer
1. Connects to Kafka broker as part of a consumer group
2. Subscribes to the specified topic
3. Kafka assigns partitions to this consumer
4. Reads messages from assigned partitions
5. Processes messages and commits offsets
6. Continues until interrupted (Ctrl+C)

### Key Features Demonstrated

1. **Message Keys**: Messages with the same key go to the same partition
2. **Offsets**: Each message has a unique offset within its partition
3. **Consumer Groups**: Multiple consumers can share the load
4. **Acknowledgment**: Messages are marked as processed
5. **Error Handling**: Retries and error logging

## Docker Management Commands

### View Kafka logs
```powershell
docker logs kafka
```

### Create a topic manually
```powershell
docker exec kafka kafka-topics --create --topic my-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
```

### List all topics
```powershell
docker exec kafka kafka-topics --list --bootstrap-server localhost:9092
```

### Describe a topic
```powershell
docker exec kafka kafka-topics --describe --topic test-topic --bootstrap-server localhost:9092
```

### Stop and remove containers
```powershell
docker stop kafka zookeeper
docker rm kafka zookeeper
```

## Troubleshooting

### Connection Refused
- Ensure Kafka container is running: `docker ps`
- Check Kafka logs: `docker logs kafka`
- Wait a few seconds after starting Kafka (it takes time to initialize)

### Topic Not Found
- Topics are auto-created by default when producer sends first message
- Or create manually using kafka-topics command above

### Consumer Not Receiving Messages
- Make sure producer ran successfully first
- Check if both are using the same topic name
- Verify broker address is correct
