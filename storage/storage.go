package storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

type RedisStorage struct {
	rdb *redis.Client
	ctx context.Context
}

func RedisConnect() *RedisStorage {

	redisStorage := new(RedisStorage)

	redisStorage.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
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

	fmt.Println("SET Key: ", key)     // debug
	fmt.Println("SET Value: ", value) // debug

	if err != nil {
		fmt.Println("Error SET in redis: ", err)
	}

	result := redisStorage.Get(key)   // debug
	fmt.Println("Try GET : ", result) // debug
}

func (redisStorage *RedisStorage) Get(key string) string {
	value, err := redisStorage.rdb.Get(redisStorage.ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}
