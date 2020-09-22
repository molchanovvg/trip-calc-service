package storage

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"os"
)

type RedisStorage struct {
	rdb *redis.Client
	ctx context.Context
}

func RedisConnect() *RedisStorage {

	redisStorage := new(RedisStorage)

	redisStorage.rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // default
		DB:       0,  // default
	})

	redisStorage.ctx = context.Background()

	pong, err := redisStorage.Ping()

	if err == nil {
		fmt.Println(pong, err)
	}
	return redisStorage
}

func (redisStorage *RedisStorage) Ping() (string, error) {

	return redisStorage.rdb.Ping(redisStorage.ctx).Result()
}

func (redisStorage *RedisStorage) Set(key string, value string) {

	err := redisStorage.rdb.Set(redisStorage.ctx, key, value, 0).Err()

	if err != nil {
		fmt.Println("Error set in Redis", key, value, err)
	}
}

func (redisStorage *RedisStorage) Get(key string) string {
	value, err := redisStorage.rdb.Get(redisStorage.ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}
