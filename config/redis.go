package config

import (
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func NewRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		PoolSize:     100,
		PoolTimeout:  time.Second,
		Addr:         os.Getenv("REDIS_HOST"),
		Password:     os.Getenv("REDIS_PASS"),
	})
}
