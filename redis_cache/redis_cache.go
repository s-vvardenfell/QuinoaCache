package redis_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background() //TODO

type RedisCache struct {
	client *redis.Client
}

func New(host, port, passw string, db int) *RedisCache {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port), //:6379
			Password: passw,
			DB:       db,
		}),
	}
}

func (r *RedisCache) set(key, value string, exp time.Duration) {
	if err := r.client.Set(
		ctx, key, value, exp*time.Second).Err(); err != nil {
		logrus.Fatalf("cannot set key-value in Redis, %v", err)
	}
}

func (r *RedisCache) get(key string) string {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		logrus.Warningf("key %s does not exist, %v", key, err)
		return ""
	} else if err != nil {
		logrus.Fatalf("cannot get value by key %s from Redis, %v", key, err)
		return ""
	} else {
		return val
	}
}

func (r *RedisCache) Create(key, value string) {
	r.set(key, value, 0)
}

func (r *RedisCache) Read(key string) string {
	return r.get((key))
}

func (r *RedisCache) Update(key, value string) {
	r.set(key, value, 0)
}

func (r *RedisCache) Delete(key string) {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		logrus.Fatalf("cannot delete key %s in Redis, %v", key, err)
	}
}
