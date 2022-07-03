#!/bin/sh
protoc --go-grpc_out=. proto/proto.proto
protoc --go_out=. proto/proto.proto

#protoc -I. --go_opt=paths=source_relative --go_out=./generated proto/service.proto
#protoc -I . proto/service.proto --go_out=. --go-grpc_out=. proto/service.proto


#path: -I./ --go_opt=paths=source_relative --go_out=./generated proto/service.proto

# protoc -I=./ --go_out=. $(find proto -type f -name '*.proto')
# protoc --go-grpc_out=. proto/service.proto