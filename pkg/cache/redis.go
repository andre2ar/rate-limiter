package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client *redis.Client
	ctx    *context.Context
}

func NewRedisClient(address string, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()

	return &RedisClient{client: client, ctx: &ctx}
}

func (c *RedisClient) Set(key string, value any, expiration time.Duration) error {
	return c.client.Set(*c.ctx, key, value, expiration).Err()
}

func (c *RedisClient) Get(key string) (any, error) {
	return c.client.Get(*c.ctx, key).Result()
}
