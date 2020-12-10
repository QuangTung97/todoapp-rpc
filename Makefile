.PHONY: generate install-tools

generate:
	@./generate.sh

install-tools:
	go install github.com/gogo/protobuf/protoc-gen-gofast
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
