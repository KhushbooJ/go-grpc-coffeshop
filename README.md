# go-grpc-coffeeshop

A minimal gRPC coffee shop demo in Go, showcasing unary and server-streaming RPC patterns.

## Overview

The project defines a `CoffeeShop` gRPC service with three RPCs:

| RPC | Type | Description |
|-----|------|-------------|
| `GetMenu` | Server-streaming | Streams menu items one batch at a time |
| `PlaceOrder` | Unary | Accepts a list of items and returns a receipt ID |
| `GetOrderStatus` | Unary | Returns the current status of an order by receipt ID |

## Project Structure

```
.
├── proto/
│   ├── coffee.proto          # Service and message definitions
│   ├── coffee.pb.go          # Generated protobuf types
│   └── coffee_grpc.pb.go     # Generated gRPC stubs
├── server/
│   └── main.go               # gRPC server (listens on :8081)
├── client/
│   └── main.go               # gRPC client demo
├── Makefile
├── go.mod
└── go.sum
```

## Prerequisites

- Go 1.21+
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Getting Started

### 1. Generate proto stubs (if needed)

```bash
make build_proto
```

### 2. Start the server

```bash
go run ./server/main.go
```

The server listens on `localhost:8081`.

### 3. Run the client

```bash
go run ./client/main.go
```

The client will:
1. Stream the full menu from the server
2. Place an order with the received items
3. Fetch and print the order status

## Dependencies

- [google.golang.org/grpc](https://pkg.go.dev/google.golang.org/grpc) v1.81.1
- [google.golang.org/protobuf](https://pkg.go.dev/google.golang.org/protobuf) v1.36.11
