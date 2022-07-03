package redis_cache

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/s-vvardenfell/Adipiscing/generated"
	"github.com/s-vvardenfell/Adipiscing/utility"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

func RunServer(cnfg map[string]interface{}) {
	grpcServ := grpc.NewServer()
	rcs := NewServer(
		cnfg["host"].(string),
		cnfg["redis_port"].(string),
		cnfg["pasword"].(string),
		cnfg["db_num"].(int),
	)
	generated.RegisterRedisCacheServiceServer(grpcServ, rcs)

	lis, err := net.Listen("tcp", ":"+cnfg["server_port"].(string))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Starting gRPC listener on port " + cnfg["server_port"].(string))
	if err := grpcServ.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Test_pckg(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	f, err := os.Open(filepath.Join(filepath.Join(wd, ".."), "resources/config_test.yml"))
	require.NoError(t, err)

	cnfg := make(map[string]interface{})

	err = yaml.Unmarshal(utility.BytesFromReader(f), &cnfg)
	require.NoError(t, err)

	t.Logf("%v", cnfg)

	go RunServer(cnfg)

	time.Sleep(1 * time.Second)

	var key, val string
	t.Log("\tGenerating random key and value")
	{
		rand.Seed(time.Now().UnixNano())
		key = strconv.Itoa(rand.Intn(1000))
		val = strconv.Itoa(rand.Intn(1000))
	}

	s := NewClientStub(cnfg["host"].(string), cnfg["server_port"].(string))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("\tSet value")
	{
		res, err := s.c.Set(ctx, &generated.Input{Key: key, Val: val, Exp: 1000})
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
