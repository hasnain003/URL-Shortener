package redisstore

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	once   sync.Once
)

func createRedisClient(address, password string) {
	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
}

func GetClient(address, password string) *redis.Client {
	once.Do(func() { createRedisClient(address, password) })
	return client
}
