package main

//go:generate go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
//go:generate go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//go:generate go run ../tools/protogen/
