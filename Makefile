
GOPATH:=$(shell go env GOPATH)
SRV_NAME:=micro-example-srv


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/example/example.proto

.PHONY: mkdir
mkdir:
	mkdir -p output/bin

.PHONY: build
build: proto mkdir
	go build -o output/bin/${SRV_NAME} main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t micro-example-srv:latest
