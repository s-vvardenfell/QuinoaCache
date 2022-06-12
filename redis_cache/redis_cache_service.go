package redis_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/s-vvardenfell/Adipiscing/generated"
)

type Server struct {
	generated.UnimplementedUserServiceServer
	client *redis.Client
}

func NewServer(host, port, passw string, db int) *Server {
	return &Server{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port), //:6379
			Password: passw,
			DB:       db,
		}),
	}
}

func (s *Server) Get(ctx context.Context, in *generated.Key) (*generated.Value, error) {
	res, err := s.client.Get(ctx, in.Key).Result()
	return &generated.Value{Val: res}, err
}

func (s *Server) Set(ctx context.Context, in *generated.Input) (*generated.Ok, error) {
	err := s.client.Set(ctx, in.Key, in.Val, time.Duration(in.Exp)*time.Second).Err()
	if err != nil {
		return &generated.Ok{Ok: false}, err
	}
	return &generated.Ok{Ok: true}, err
}
