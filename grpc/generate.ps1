# PowerShell script to generate gRPC code from proto files

Write-Host "Generating gRPC code from proto files..." -ForegroundColor Green

# Check if protoc is installed
if (-not (Get-Command protoc -ErrorAction SilentlyContinue)) {
    Write-Host "Error: protoc is not installed!" -ForegroundColor Red
    Write-Host "Please install protoc from: https://grpc.io/docs/protoc-installation/" -ForegroundColor Yellow
    exit 1
}

# Generate Go code
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/calculator.proto

if ($LASTEXITCODE -eq 0) {
    Write-Host "✅ Code generation successful!" -ForegroundColor Green
    Write-Host "Generated files:" -ForegroundColor Cyan
    Get-ChildItem -Path "proto" -Filter "*.pb.go" | ForEach-Object { Write-Host "  - $($_.Name)" -ForegroundColor Cyan }
} else {
    Write-Host "❌ Code generation failed!" -ForegroundColor Red
    exit 1
}
