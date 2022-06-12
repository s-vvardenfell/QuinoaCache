package redis_cache

import (
	"log"

	"github.com/s-vvardenfell/Adipiscing/generated"
	"google.golang.org/grpc"
)

type RedisClientStub struct {
	c generated.UserServiceClient
}

func NewClientStub(host, port string) *RedisClientStub {
	conn, err := grpc.Dial(host+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &RedisClientStub{
		c: generated.NewUserServiceClient(conn),
	}
}
