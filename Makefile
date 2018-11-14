
GOPATH:=$(shell go env GOPATH)
SRV_NAME:=micro-rpc
API_NAME:=micro-api
PROTO_FILE:=proto/business/business.proto


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. ${PROTO_FILE}

.PHONY: clean
clean:
	rm -rf output/

.PHONY: mkdir
mkdir:
	mkdir -p output/bin

.PHONY: api
api:
	go build -o output/bin/${API_NAME} main_api.go plugin.go

.PHONY: rpc
rpc:
	go build -o output/bin/${SRV_NAME} main_rpc.go plugin.go

.PHONY: build
build: clean proto mkdir api rpc

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t me-moye-srv:latest
