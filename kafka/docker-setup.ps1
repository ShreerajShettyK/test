# PowerShell script to set up Kafka with Docker

Write-Host "=== Kafka Docker Setup ===" -ForegroundColor Green
Write-Host ""

# Check if Docker is running
Write-Host "Checking if Docker is running..." -ForegroundColor Yellow
docker version > $null 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Docker is not running. Please start Docker Desktop." -ForegroundColor Red
    exit 1
}
Write-Host "Docker is running!" -ForegroundColor Green
Write-Host ""

# Create network if it doesn't exist
Write-Host "Creating Docker network 'kafka-network'..." -ForegroundColor Yellow
docker network create kafka-network 2>$null
if ($LASTEXITCODE -eq 0) {
    Write-Host "Network created successfully!" -ForegroundColor Green
} else {
    Write-Host "Network already exists or created." -ForegroundColor Cyan
}
Write-Host ""

# # Stop and remove existing containers if they exist
# Write-Host "Cleaning up existing containers..." -ForegroundColor Yellow
# docker stop kafka zookeeper 2>$null
# docker rm kafka zookeeper 2>$null
# Write-Host "Cleanup complete!" -ForegroundColor Green
# Write-Host ""

# # Start Zookeeper
# Write-Host "Starting Zookeeper..." -ForegroundColor Yellow
# docker run -d `
#     --name zookeeper `
#     --network kafka-network `
#     -p 2181:2181 `
#     -e ZOOKEEPER_CLIENT_PORT=2181 `
#     -e ZOOKEEPER_TICK_TIME=2000 `
#     confluentinc/cp-zookeeper:latest

# if ($LASTEXITCODE -eq 0) {
#     Write-Host "Zookeeper started successfully!" -ForegroundColor Green
# } else {
#     Write-Host "Error starting Zookeeper!" -ForegroundColor Red
#     exit 1
# }
# Write-Host ""

# # Wait for Zookeeper to be ready
# Write-Host "Waiting for Zookeeper to be ready..." -ForegroundColor Yellow
# Start-Sleep -Seconds 5

# # Start Kafka
# Write-Host "Starting Kafka..." -ForegroundColor Yellow
# docker run -d `
#     --name kafka `
#     --network kafka-network `
#     -p 9092:9092 `
#     -e KAFKA_BROKER_ID=1 `
#     -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 `
#     -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT `
#     -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:29092,PLAINTEXT_HOST://localhost:9092 `
#     -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:29092,PLAINTEXT_HOST://0.0.0.0:9092 `
#     -e KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT `
#     -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 `
#     -e KAFKA_TRANSACTION_STATE_LOG_MIN_ISR=1 `
#     -e KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR=1 `
#     -e KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS=0 `
#     -e KAFKA_AUTO_CREATE_TOPICS_ENABLE=true `
#     confluentinc/cp-kafka:7.5.0

# if ($LASTEXITCODE -eq 0) {
#     Write-Host "Kafka started successfully!" -ForegroundColor Green
# } else {
#     Write-Host "Error starting Kafka!" -ForegroundColor Red
#     exit 1
# }
# Write-Host ""

# # Wait for Kafka to be ready
# Write-Host "Waiting for Kafka to be ready (this may take 30-40 seconds)..." -ForegroundColor Yellow
# Start-Sleep -Seconds 30

# Check container status
Write-Host "Checking container status..." -ForegroundColor Yellow
docker ps --filter "name=kafka" --filter "name=zookeeper"
Write-Host ""

Write-Host "=== Setup Complete! ===" -ForegroundColor Green
Write-Host ""
Write-Host "Kafka is running on: localhost:9092" -ForegroundColor Cyan
Write-Host "Zookeeper is running on: localhost:2181" -ForegroundColor Cyan
Write-Host ""
Write-Host "You can now run the producer and consumer:" -ForegroundColor Yellow
Write-Host "  Producer: go run . -mode producer -topic test-topic" -ForegroundColor White
Write-Host "  Consumer: go run . -mode consumer -topic test-topic" -ForegroundColor White
Write-Host ""
Write-Host "To view logs:" -ForegroundColor Yellow
Write-Host "  docker logs kafka" -ForegroundColor White
Write-Host "  docker logs zookeeper" -ForegroundColor White
Write-Host ""
Write-Host "To stop:" -ForegroundColor Yellow
Write-Host "  docker stop kafka zookeeper" -ForegroundColor White
