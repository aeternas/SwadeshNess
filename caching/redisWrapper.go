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
	return "", nil
}

func (rcw *redisCachingWrapper) SaveCachedValue(k, v string) error {
	return nil
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
	// Output: PONG <nil>
}
