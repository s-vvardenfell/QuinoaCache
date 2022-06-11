#!/bin/sh
protoc --go-grpc_out=. proto/service.proto
protoc --go_out=. proto/service.proto
