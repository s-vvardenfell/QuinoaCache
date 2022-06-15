#!/bin/sh
protoc --go-grpc_out=. proto/service.proto
protoc --go_out=. proto/service.proto

#protoc -I. --go_opt=paths=source_relative --go_out=./generated proto/service.proto
#protoc -I . proto/service.proto --go_out=. --go-grpc_out=. proto/service.proto