package config

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	REDIS_URL := os.Getenv("REDIS_URL")
	PASSWORD := os.Getenv("PASSWORD")
	DB := os.Getenv("DB")
	DBi, err := strconv.Atoi(DB)
	if err != nil {
		panic("not connected")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: PASSWORD,
		DB:       DBi,
	})

	return &RedisClient{Client: client}
}
