package redis_cache

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/s-vvardenfell/Adipiscing/generated"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

const (
	redisPort = "6379"
	servePort = ":50051"
	host      = "localhost"
)

func RunServer() {
	grpcServ := grpc.NewServer()
	rcs := NewServer("localhost", redisPort, "", 0)
	generated.RegisterUserServiceServer(grpcServ, rcs)

	lis, err := net.Listen("tcp", servePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Starting gRPC listener on port " + servePort)
	if err := grpcServ.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Test_pckg(t *testing.T) {

	go RunServer()

	time.Sleep(2 * time.Second)

	var key, val string
	t.Log("\tGenerating random key and value")
	{
		rand.Seed(time.Now().UnixNano())
		key = strconv.Itoa(rand.Intn(1000))
		val = strconv.Itoa(rand.Intn(1000))
	}

	s := NewClientStub(host, servePort)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("\tSet value")
	{
		res, err := s.c.Set(ctx, &generated.Input{Key: key, Val: val, Exp: 0})
		require.NoError(t, err)
		require.Equal(t, res.Ok, true)
	}

	t.Log("\tGet value")
	{
		res, err := s.c.Get(ctx, &generated.Key{Key: key})
		require.NoError(t, err)
		require.Equal(t, res.Val, val)
	}

	t.Log("\tGenerating new random key and value")
	{
		rand.Seed(time.Now().UnixNano())
		key = strconv.Itoa(rand.Intn(1000))
		val = strconv.Itoa(rand.Intn(1000))
	}

	t.Log("\tSet value with expire time 5 seconds")
	{
		res, err := s.c.Set(ctx, &generated.Input{Key: key, Val: val, Exp: 5})
		require.NoError(t, err)
		require.Equal(t, res.Ok, true)
	}

	t.Log("\tGet value with expire time after 1 second while it exists")
	{
		time.Sleep(1 * time.Second)
		res, err := s.c.Get(ctx, &generated.Key{Key: key})
		require.NoError(t, err)
		require.Equal(t, res.Val, val)
	}

	t.Log("\tGet value with expire time after 5 seconds when it does not exists")
	{
		time.Sleep(4 * time.Second)
		_, err := s.c.Get(ctx, &generated.Key{Key: key})
		require.Error(t, err)
	}
}
