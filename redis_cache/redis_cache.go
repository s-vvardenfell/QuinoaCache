package redis_cache

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/s-vvardenfell/Adipiscing/utility"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background() //TODO

type RedisCache struct {
	client  *redis.Client
	expires time.Duration
}

//TODO change to config.host etc
func New(host string, db int, exp time.Duration) *RedisCache {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:6379", host), //TODO mb port from config
			Password: "",                           //TODO from config
			DB:       db,                           //TODO from config
		}),
		expires: exp,
	}
}

func (r *RedisCache) Set(key string, value io.Reader) {
	if err := r.client.Set(
		ctx, key, utility.BytesFromReader(value), r.expires*time.Second).Err(); err != nil {
		logrus.Fatalf("cannot set key-value in Redis, %v", err)
	}
}

func (r *RedisCache) Get(key string) []byte {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		logrus.Warning("key %s does not exist, %v", key, err)
		return nil
	} else if err != nil {
		logrus.Fatalf("cannot get value by key %s from Redis, %v", key, err)
	} else {
		return []byte(val)
	}
	return nil
}

func (r *RedisCache) Create() {

}

func (r *RedisCache) Read() {

}

func (r *RedisCache) Update() {

}

func (r *RedisCache) Delete() {

}
