# Quick Start Guide

## ✅ Setup Complete!

Your Kafka environment is ready! Here's what's running:
- **Kafka Broker**: localhost:9092
- **Zookeeper**: localhost:2181
- **Docker Containers**: kafka, zookeeper

## 🚀 Quick Test (Run these in order)

### Step 1: Open a NEW PowerShell terminal and run the Consumer
```powershell
cd c:\Users\Relanto\Desktop\relanto-cisco\test\kafka
.\kafka-demo.exe -mode consumer -topic test-topic
```
This will start listening for messages. Keep this terminal open!

### Step 2: In THIS terminal, run the Producer
```powershell
.\kafka-demo.exe -mode producer -topic test-topic
```
This will send 5 messages to the topic.

### Step 3: Check the Consumer terminal
You should see the consumer receiving and displaying all 5 messages!

## 📊 What You Just Saw

### Producer Output
- Shows each message being sent
- Displays the partition and offset assigned
- Messages with the same key go to the same partition

### Consumer Output
- Shows each message being received
- Displays message metadata (partition, offset, timestamp)
- Messages are consumed in order within a partition

## 🎯 Understanding Kafka Topics

### Key Concepts Demonstrated:

1. **Topic** (`test-topic`): A named stream where messages are stored
2. **Partition** (0): Topics are divided into partitions for scalability
3. **Offset** (0,1,2,3,4): Each message has a unique position in its partition
4. **Key** (user-1, user-2, etc.): Used to determine which partition a message goes to
5. **Consumer Group**: Multiple consumers can share the load

### Message Flow:
```
Producer → [test-topic] → Consumer
             ├── Partition 0
             │   ├── Offset 0: user-1 message
             │   ├── Offset 1: user-2 message
             │   ├── Offset 2: user-3 message
             │   ├── Offset 3: user-1 message
             │   └── Offset 4: user-4 message
```

## 🧪 Experiments to Try

### 1. Run Multiple Consumers
Open 2-3 terminals and run the consumer in each. Kafka will distribute partitions among them!
```powershell
.\kafka-demo.exe -mode consumer -topic test-topic -group my-group
```

### 2. Send Custom Messages
Modify `producer.go` to send your own messages

### 3. Check Topic Info
```powershell
docker exec kafka kafka-topics --describe --topic test-topic --bootstrap-server localhost:9092
```

### 4. Create a New Topic
```powershell
docker exec kafka kafka-topics --create --topic my-new-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
```

### 5. List All Topics
```powershell
docker exec kafka kafka-topics --list --bootstrap-server localhost:9092
```

### 6. View Kafka Logs
```powershell
docker logs kafka
```

## 🛠️ Useful Commands

### Build the application
```powershell
go build -o kafka-demo.exe
```

### Run Producer
```powershell
.\kafka-demo.exe -mode producer -topic test-topic
```

### Run Consumer
```powershell
.\kafka-demo.exe -mode consumer -topic test-topic -group test-group
```

### Different Topic
```powershell
.\kafka-demo.exe -mode producer -topic my-topic
.\kafka-demo.exe -mode consumer -topic my-topic
```

## 🔍 Troubleshooting

### If messages aren't being received:
1. Make sure consumer started before producer
2. Check both are using the same topic name
3. Verify Kafka is running: `docker ps`

### If connection fails:
1. Wait 30-40 seconds after starting containers
2. Check logs: `docker logs kafka`
3. Restart: `.\docker-setup.ps1`

## 🛑 Stop Everything

### Stop consuming (in consumer terminal)
Press `Ctrl+C`

### Stop Docker containers
```powershell
docker stop kafka zookeeper
```

### Remove containers
```powershell
docker rm kafka zookeeper
```

### Restart everything
```powershell
.\docker-setup.ps1
```

## 📚 Learn More

Check `README.md` for detailed explanations of:
- How Kafka topics work
- Producer and Consumer architecture
- Code walkthrough
- Advanced configuration options

---

**Happy Kafka Learning! 🎉**
