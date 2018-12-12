package caching

import (
	"fmt"
	Redis "github.com/go-redis/redis"
)

type RedisCachingWrapper interface {
	GetCachedValue(k string) (string, error)
	SaveCachedValue(k, v string) error
}

type redisCachingWrapper struct {
	RedisClient *Redis.Client
}

func NewRedisCachingWrapper() RedisCachingWrapper {
	var redisClient *Redis.Client = ExampleNewClient()
	return &redisCachingWrapper{RedisClient: redisClient}
}

func (rcw *redisCachingWrapper) GetCachedValue(k string) (string, error) {
	val, err := rcw.RedisClient.Get(k).Result()
	if err != nil {
		return "", err
	}

	//if err == redis.Nil {
	return val, nil
}

func (rcw *redisCachingWrapper) SaveCachedValue(k, v string) error {
	err := rcw.RedisClient.Set(k, v, 0).Err()
	return err
}

func ExampleNewClient() *Redis.Client {
	client := Redis.NewClient(&Redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
