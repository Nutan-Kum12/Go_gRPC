# Go gRPC Streaming Examples

A comprehensive gRPC implementation in Go demonstrating all four types of RPC communication patterns with practical examples.

## 📋 Table of Contents
- [Overview](#overview)
- [Project Structure](#project-structure)
- [RPC Types Explained](#rpc-types-explained)
- [Real-World Use Cases](#real-world-use-cases)
- [Setup & Installation](#setup--installation)
- [Usage](#usage)
- [File Explanations](#file-explanations)
- [Flow Diagrams](#flow-diagrams)
- [Common Issues](#common-issues)

## 🎯 Overview

This project demonstrates all four gRPC communication patterns:
1. **Unary RPC** - Simple request/response
2. **Server Streaming** - One request, multiple responses
3. **Client Streaming** - Multiple requests, one response
4. **Bidirectional Streaming** - Multiple requests and responses simultaneously

## 📁 Project Structure

```
Go_gRPC/
├── proto/                     # Protocol Buffer definitions
│   ├── greet.proto           # Service and message definitions
│   ├── greet.pb.go           # Generated Go structs
│   └── greet_grpc.pb.go      # Generated gRPC client/server code
├── server/                   # Server implementation
│   ├── main.go              # Server setup and startup
│   ├── unary.go             # Unary RPC implementation
│   ├── server_stream.go     # Server streaming implementation
│   ├── client_stream.go     # Client streaming implementation
│   └── bi_stream.go         # Bidirectional streaming implementation
├── client/                   # Client implementation
│   ├── main.go              # Client connection setup
│   ├── unary.go             # Unary RPC client calls
│   ├── server_stream.go     # Server streaming client calls
│   ├── client_stream.go     # Client streaming client calls
│   └── bi_stream.go         # Bidirectional streaming client calls
├── makefile                 # Build and run commands
├── go.mod                   # Go module dependencies
└── README.md               # This file
```

## 🔄 RPC Types Explained

### 1. Unary RPC (Request → Response)
**What**: Client sends one request, server sends one response.
```
Client ────[Request]────> Server
Client <───[Response]──── Server
```

**Implementation**:
- **Server**: `server/unary.go` - `SayHello()`
- **Client**: `client/unary.go` - `callSayHello()`

**Code Flow**:
```go
// Client sends NoParam{}, Server responds with HelloResponse
req := &pb.NoParam{}
res, err := client.SayHello(ctx, req)
```

### 2. Server Streaming (Request → Stream of Responses)
**What**: Client sends one request, server sends multiple responses over time.
```
Client ────[Request]────────> Server
Client <───[Response 1]────── Server
Client <───[Response 2]────── Server
Client <───[Response 3]────── Server
Client <───[EOF]──────────── Server
```

**Implementation**:
- **Server**: `server/server_stream.go` - `SayHelloServerStreaming()`
- **Client**: `client/server_stream.go` - `callSayHelloServerStream()`

**Code Flow**:
```go
// Client sends list of names, Server streams back greetings
names := &pb.NameList{Names: []string{"Alice", "Bob", "Charlie"}}
stream, err := client.SayHelloServerStreaming(ctx, names)
for {
    resp, err := stream.Recv() // Receive each response
    if err == io.EOF { break } // Server finished
}
```

### 3. Client Streaming (Stream of Requests → Response)
**What**: Client sends multiple requests over time, server sends one final response.
```
Client ────[Request 1]────> Server
Client ────[Request 2]────> Server  
Client ────[Request 3]────> Server
Client ────[EOF]──────────> Server
Client <───[Final Response] Server
```

**Implementation**:
- **Server**: `server/client_stream.go` - `SayHelloClientStreaming()`
- **Client**: `client/client_stream.go` - `callSayHelloClientStream()`

**Code Flow**:
```go
// Client streams multiple names, Server responds with collected messages
stream, err := client.SayHelloClientStreaming(ctx)
for _, name := range names {
    stream.Send(&pb.HelloRequest{Name: name})
}
resp, err := stream.CloseAndRecv() // Get final response
```

### 4. Bidirectional Streaming (Stream ↔ Stream)
**What**: Both client and server can send multiple messages simultaneously.
```
Client ────[Request 1]────> Server
Client <───[Response 1]──── Server
Client ────[Request 2]────> Server
Client <───[Response 2]──── Server
Client ────[EOF]──────────> Server
Client <───[EOF]──────────── Server
```

**Implementation**:
- **Server**: `server/bi_stream.go` - `SayHelloBidirectionalStreaming()`
- **Client**: `client/bi_stream.go` - `callSayHelloBidirectionalStreaming()`

**Code Flow**:
```go
// Both send and receive simultaneously using goroutines
stream, err := client.SayHelloBidirectionalStreaming(ctx)

// Goroutine for receiving
go func() {
    for {
        resp, err := stream.Recv()
        if err == io.EOF { break }
        log.Println(resp.Message)
    }
    close(waitc) // Signal completion
}()

// Send requests
for _, name := range names {
    stream.Send(&pb.HelloRequest{Name: name})
}
stream.CloseSend()
<-waitc // Wait for receiving to complete
```

## 🌍 Real-World Use Cases

### Unary RPC
- **Authentication**: Login with username/password
- **User Profile**: Get user details by ID
- **Payment Processing**: Process single transaction
- **Configuration**: Get application settings

### Server Streaming
- **Live Updates**: Stock prices, sports scores
- **File Download**: Large file broken into chunks
- **Chat Messages**: Receiving messages in a chat room
- **Log Streaming**: Real-time application logs
- **News Feed**: Continuous feed of articles

### Client Streaming
- **File Upload**: Large file sent in chunks
- **Sensor Data**: IoT devices sending batch readings
- **Bulk Operations**: Multiple database inserts
- **Analytics**: Sending multiple events for processing

### Bidirectional Streaming
- **Chat Applications**: Real-time messaging
- **Collaborative Editing**: Google Docs-style editing
- **Gaming**: Real-time multiplayer games
- **Video Calls**: Audio/video streaming
- **Trading Platforms**: Real-time order processing

## 🚀 Setup & Installation

### Prerequisites
- Go 1.19+ installed
- Protocol Buffers compiler (`protoc`)
- gRPC Go plugins

### Install Dependencies
```bash
# Install Protocol Buffer compiler
# Windows (using chocolatey):
choco install protoc

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install project dependencies
go mod tidy
```

### Generate Protocol Buffer Code
```bash
make proto
# Or manually:
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

## 🎮 Usage

### Run Individual Components

1. **Start the Server**:
```bash
make run
# Server starts on localhost:8080
```

2. **Run Client** (in another terminal):
```bash
make client
# Executes the client with current configuration
```

3. **Run Both Simultaneously**:
```bash
make run-all
# Opens server and client in separate windows
```

### Switch Between RPC Types

Edit `client/main.go` to test different RPC patterns:

```go
func main() {
    // ... connection setup ...
    client := pb.NewGreetServiceClient(conn)
    names := &pb.NameList{Names: []string{"Alice", "Bob", "Charlie"}}
    
    // Choose one:
    callSayHello(client)                              // Unary
    callSayHelloServerStream(client, names)           // Server Streaming
    callSayHelloClientStream(client, names)           // Client Streaming
    callSayHelloBidirectionalStreaming(client, names) // Bidirectional
}
```

## 📄 File Explanations

### Protocol Buffer Files

**`proto/greet.proto`**
- Defines the gRPC service interface
- Specifies all 4 RPC methods
- Defines message structures (requests/responses)

**`proto/greet.pb.go`** & **`proto/greet_grpc.pb.go`**
- Auto-generated Go code from `.proto` file
- Contains structs, interfaces, and gRPC client/server code
- **Never edit manually** - regenerate with `make proto`

### Server Files

**`server/main.go`**
- Server initialization and startup
- Creates gRPC server on port 8080
- Registers the GreetService implementation

**`server/unary.go`**
```go
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error)
```
- Simple request/response handler
- Returns single greeting message

**`server/server_stream.go`**
```go
func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error
```
- Receives list of names
- Streams back individual greetings with delays
- Uses `stream.Send()` to send multiple responses

**`server/client_stream.go`**
```go
func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error
```
- Receives stream of names from client
- Collects all names and sends back aggregated response
- Uses `stream.Recv()` in loop until `io.EOF`

**`server/bi_stream.go`**
```go
func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error
```
- Handles simultaneous sending and receiving
- Uses goroutines for concurrent processing
- Processes each request immediately and responds

### Client Files

**`client/main.go`**
- Client connection setup
- Creates gRPC client connection to server
- Calls the desired RPC method

**`client/unary.go`**
- Simple RPC call with timeout
- Demonstrates basic error handling

**`client/server_stream.go`**
- Initiates server streaming
- Loops to receive all responses until `io.EOF`
- Handles timeout appropriately for streaming

**`client/client_stream.go`**
- Sends multiple requests over stream
- Uses `CloseAndRecv()` to finish and get response

**`client/bi_stream.go`**
- Most complex client implementation
- Uses goroutines for concurrent send/receive
- Uses channels for synchronization (`waitc`)

## 🔄 Flow Diagrams

### Unary RPC Flow
```
Client                    Server
  |                         |
  |-----> SayHello() ------>|
  |                         | (Process)
  |<----- Response ---------|
  |                         |
```

### Server Streaming Flow
```
Client                    Server
  |                         |
  |----> NameList --------->|
  |                         | (Start streaming)
  |<---- Response 1 --------|
  |<---- Response 2 --------|
  |<---- Response 3 --------|
  |<---- EOF ---------------|
```

### Client Streaming Flow
```
Client                    Server
  |                         |
  |----> Request 1 -------->| (Collect)
  |----> Request 2 -------->| (Collect)
  |----> Request 3 -------->| (Collect)
  |----> EOF -------------->| (Process all)
  |<---- Final Response ----|
```

### Bidirectional Streaming Flow
```
Client                    Server
  |                         |
  |----> Request 1 -------->|----> Response 1
  |<---- Response 1 --------|
  |----> Request 2 -------->|----> Response 2
  |<---- Response 2 --------|
  |----> EOF -------------->|
  |<---- EOF ---------------|

```

## 🎯 Key Learning Points

1. **Unary**: Best for simple request/response operations
2. **Server Streaming**: Use when server needs to send continuous updates
3. **Client Streaming**: Use when client needs to send large amounts of data
4. **Bidirectional**: Use for real-time interactive applications
5. **Channels**: Essential for synchronization in bidirectional streaming
6. **Context**: Important for timeouts and cancellation
7. **Error Handling**: Always check for `io.EOF` in streaming operations
