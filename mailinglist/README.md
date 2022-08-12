# MailingList - Microservice

## Description

## Purpose

## Product functions (brief description)

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
```
go mod tidy
go run ./pixl
```

[//]: # (For this project, we'll be creating a microservice that manages an email list, will be creating a)

[//]: # ()
[//]: # (JSON and gRPC API server will also be creating a gRPC API client to communicate with the server or utilizing)

[//]: # ()
[//]: # (protocol buffers will have the database backed by SQLite, will utilize good routines to run both servers)

[//]: # ()
[//]: # (in a single application and it's going to be a CLI project)


