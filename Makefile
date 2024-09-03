GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. define/AppAuthServer.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. define/UserManager.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. define/SignatureEncryption.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. define/BackendAPI.proto