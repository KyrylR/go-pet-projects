# MailingList - Microservice

Microservice that provides the management of an email address list.

## Description

Created a microservice that hosts two APIs - the gRPC API and the JSON API.

## Purpose

The main goal of this work is to learn how to work with APIs and microservices 
as well as to consolidate the knowledge gained.

## Product functions (brief description)

* JSON and gRPC API servers.
* Using goroutines to run both servers in the same application.
* Using gRPC client API to communicate with the server using protocol buffers.
* SQLite database integration.


## Dependencies

This project requires:
* `gcc` compiler installed 
* `protobuf` code generation tools
* `go` version 1.9

### Install protobuf compiler

Install the `protoc` tool using the instructions available at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).

Alternatively you can download a pre-built binary from [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases) and placing the extracted binary somewhere in your `$PATH`.

### Install Go protobuf codegen tools

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

### Generate Go code from .proto files

```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  Proto/mail.proto
```

## Executing program

The program is invoked according to the scheme:

* Server
```
go mod tidy
go run ./server
```

* Client
```
go run ./client
```
