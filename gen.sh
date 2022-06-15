#!/bin/sh
$GOPATH/bin/protoc --go-grpc_out=. proto/service.proto
$GOPATH/bin/protoc --go_out=. proto/service.proto

#protoc -I . proto/service.proto --go_out=. --go-grpc_out=. proto/service.proto