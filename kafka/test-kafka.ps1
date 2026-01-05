# PowerShell script to test Kafka producer and consumer

Write-Host "=== Kafka Producer and Consumer Test ===" -ForegroundColor Green
Write-Host ""

Write-Host "This script will:" -ForegroundColor Yellow
Write-Host "  1. Start the consumer in the background" -ForegroundColor White
Write-Host "  2. Wait a few seconds" -ForegroundColor White
Write-Host "  3. Run the producer to send messages" -ForegroundColor White
Write-Host "  4. Show you the consumer output" -ForegroundColor White
Write-Host ""

# Start consumer in a new PowerShell window
Write-Host "Starting Consumer in a new window..." -ForegroundColor Yellow
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PWD'; Write-Host 'Consumer Window - Press Ctrl+C to stop' -ForegroundColor Green; go run . -mode consumer -topic test-topic"

Write-Host "Consumer started! Wait 5 seconds for it to initialize..." -ForegroundColor Green
Start-Sleep -Seconds 5

# Run producer
Write-Host ""
Write-Host "Running Producer to send messages..." -ForegroundColor Yellow
Write-Host ""
go run . -mode producer -topic test-topic

Write-Host ""
Write-Host "=== Test Complete! ===" -ForegroundColor Green
Write-Host ""
Write-Host "Check the Consumer window to see the received messages!" -ForegroundColor Cyan
Write-Host "Press Ctrl+C in the Consumer window to stop it." -ForegroundColor Yellow
