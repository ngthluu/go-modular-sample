package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ Cache = (*RedisCache)(nil)

var ctx = context.Background()

type RedisCache struct {
	Conn *redis.Client
}

func NewRedisCache() *RedisCache {
	addr := fmt.Sprintf("%s:%s", os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return &RedisCache{Conn: rdb}
}

func (c *RedisCache) Set(key string, value string, ttl int) error {
	return c.Conn.Set(ctx, key, value, time.Duration(ttl)*time.Minute).Err()
}

func (c *RedisCache) Update(key string, value string) error {
	return c.Conn.SetArgs(ctx, key, value, redis.SetArgs{KeepTTL: true}).Err()
}

func (c *RedisCache) Get(key string) (string, error) {
	val, err := c.Conn.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *RedisCache) Remove(key string) error {
	_, err := c.Conn.Del(ctx, key).Result()
	return err
}

func (c *RedisCache) Close() {
	c.Conn.Close()
}
