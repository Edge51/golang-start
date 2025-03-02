// internal/cache/redis.go
package cache

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
    client *redis.Client
}

func New() *RedisClient {
    return &RedisClient{
        client: redis.NewClient(&redis.Options{
            Addr: "localhost:6379", // 硬编码先测试
        }),
    }
}

func (r *RedisClient) SetURL(key, value string) error {
    return r.client.Set(ctx, key, value, 0).Err()
}