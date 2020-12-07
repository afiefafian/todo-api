package entity

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Key   string
	Value string
	TTL   int64
}

// CacheRepository represent the cache's repository contract
type CacheRepository interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, c *Cache) error
	SetWithTTL(ctx context.Context, c *Cache) error
	Delete(ctx context.Context, key string) error
	TTL(ctx context.Context, key string) int
	Decrement(ctx context.Context, key string) error
}
