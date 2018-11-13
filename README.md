# Micro-Example Service

This is the Micro-Example service

Generated with

```
micro new github.com/rockdragon/micro-example --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.micro-example
- Type: srv
- Alias: micro-example

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install protobuf
brew install protobuf

# install protoc-gen-go
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

# install protoc-gen-micro
go get -u github.com/micro/protoc-gen-micro

# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
# run the RPC
./boot_rpc.sh

# run the API
./boot_api.sh

# run the API proxy
micro api
```

Build a docker image
```
make docker
```

