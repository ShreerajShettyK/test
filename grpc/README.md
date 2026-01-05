# gRPC Calculator Example

This is a basic gRPC example that demonstrates:
- Defining a service using Protocol Buffers
- Implementing a gRPC server
- Creating a gRPC client
- Making RPC calls

## Structure

- `proto/` - Protocol Buffer definitions
- `server/` - Server implementation
- `client/` - Client implementation

## Steps to Run

1. Generate code from proto files:
   ```powershell
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/calculator.proto
   ```

2. Run the server:
   ```powershell
   go run server/main.go
   ```

3. Run the client (in another terminal):
   ```powershell
   go run client/main.go
   ```

## How gRPC Works

1. **Define Service** - Write .proto file defining service methods and messages
2. **Generate Code** - Use protoc compiler to generate client/server code
3. **Implement Server** - Write business logic implementing the service interface
4. **Create Client** - Use generated client stub to call remote methods
5. **Communication** - Client and server communicate over HTTP/2 using protobuf
