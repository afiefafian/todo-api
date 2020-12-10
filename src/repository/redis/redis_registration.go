package redis

import (
	"context"
	"fmt"
	"github.com/afiefafian/todo-api/src/entity"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisRegistrationRepository struct {
	memDB     *redis.Client
	keyPrefix string
}

// NewRedisRegistrationRepository create an data to represent cache.Repository interface
func NewRedisRegistrationRepository(memDB *redis.Client) entity.CacheRepository {
	return &redisRegistrationRepository{memDB, "register"}
}

func (r redisRegistrationRepository) Get(ctx context.Context, key string) *redis.StringCmd {
	formattedKey := fmt.Sprintf(`%s:%s`, r.keyPrefix, key)
	return r.memDB.Get(ctx, formattedKey)
}

func (r redisRegistrationRepository) Set(ctx context.Context, c *entity.Cache) error {
	key := fmt.Sprintf(`%s:%s`, r.keyPrefix, c.Key)
	r.memDB.SetEX(ctx, key, c.Value, 0)
	return nil
}

func (r redisRegistrationRepository) SetWithTTL(ctx context.Context, c *entity.Cache) error {
	key := fmt.Sprintf(`%s:%s`, r.keyPrefix, c.Key)
	ttl := time.Duration(c.TTL) * time.Second

	r.memDB.SetEX(ctx, key, c.Value, ttl)
	return nil
}

func (r redisRegistrationRepository) Delete(ctx context.Context, key string) error {
	formattedKey := fmt.Sprintf(`%s:%s`, r.keyPrefix, key)
	keys := r.memDB.Keys(ctx, formattedKey)

	for _, v := range keys.Val() {
		r.memDB.Del(ctx, v)
	}

	return nil
}

func (r redisRegistrationRepository) Decrement(ctx context.Context, key string) error {
	formattedKey := fmt.Sprintf(`%s:%s`, r.keyPrefix, key)
	r.memDB.Decr(ctx, formattedKey)
	return nil
}

func (r redisRegistrationRepository) TTL(ctx context.Context, key string) int {
	formattedKey := fmt.Sprintf(`%s:%s`, r.keyPrefix, key)
	ttl := r.memDB.TTL(ctx, formattedKey)
	return int(ttl.Val())
}
