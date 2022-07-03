package redis_cache

import (
	"fmt"

	"github.com/s-vvardenfell/QuinoaCache/generated"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RedisClientStub struct {
	c generated.RedisCacheServiceClient
}

func NewClientStub(host, port string) *RedisClientStub {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("cannot connect to host <%s> and port <%s>: %v", host, port, err)
	}
	return &RedisClientStub{
		c: generated.NewRedisCacheServiceClient(conn),
	}
}
