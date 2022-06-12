package redis_cache

import (
	"fmt"

	"github.com/s-vvardenfell/Adipiscing/generated"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type RedisClientStub struct {
	c generated.UserServiceClient
}

func NewClientStub(host, port string) *RedisClientStub {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("cannot connect to host< %s> and port <%s>: %v", host, port, err)
	}
	return &RedisClientStub{
		c: generated.NewUserServiceClient(conn),
	}
}
