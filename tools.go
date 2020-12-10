// +build tools

package tools

import (
	_ "github.com/gogo/protobuf/protoc-gen-gofast"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)
