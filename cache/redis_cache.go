package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Get(key string) (string, error) {
	ctx := context.Background()
	return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Set(key string, value string, expiration time.Duration) {
	ctx := context.Background()
	r.client.Set(ctx, key, value, expiration)
}
