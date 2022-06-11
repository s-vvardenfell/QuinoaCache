package redis_cache

import (
	"context"
	"time"

	"github.com/s-vvardenfell/Adipiscing/generated"
)

type Server struct {
	generated.UnimplementedUserServiceServer
	Rc *RedisCache
}

func (s *Server) Get(ctx context.Context, in *generated.Key) (*generated.Value, error) {
	res, err := s.Rc.get(in.Key)
	return &generated.Value{Val: res}, err
}

func (s *Server) Set(ctx context.Context, in *generated.Input) (*generated.Ok, error) {
	err := s.Rc.set(in.Key, in.Val, time.Duration(in.Exp))
	if err != nil {
		return &generated.Ok{Ok: false}, err
	}
	return &generated.Ok{Ok: true}, err
}
