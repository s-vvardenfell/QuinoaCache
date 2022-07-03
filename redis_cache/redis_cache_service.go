package redis_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/s-vvardenfell/QuinoaCache/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	generated.UnimplementedRedisCacheServiceServer
	client *redis.Client
}

func NewServer(host, port, passw string, db int) *Server {
	return &Server{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: passw,
			DB:       db,
		}),
	}
}

func (s *Server) Get(ctx context.Context, in *generated.Key) (*generated.Value, error) {
	res, err := s.client.Get(ctx, in.Key).Result()
	if err == redis.Nil {
		return nil, status.Errorf(codes.Internal,
			"key %s does not exist, %v", in.Key, err)
	} else if err != nil {
		return nil, status.Errorf(codes.Internal,
			"cannot get value by key %s from Redis, %v", in.Key, err)
	} else {
		return &generated.Value{Val: res}, nil
	}

}

func (s *Server) Set(ctx context.Context, in *generated.Input) (*generated.Ok, error) {
	err := s.client.Set(ctx, in.Key, in.Val, time.Duration(in.Exp)*time.Second).Err()
	if err != nil {
		return &generated.Ok{Ok: false}, status.Errorf(codes.Internal,
			"cannot set value  %s with key %s to Redis, %v", in.Val, in.Key, err)
	}
	return &generated.Ok{Ok: true}, nil
}
