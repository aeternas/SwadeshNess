package caching

import (
	"fmt"
	Redis "github.com/go-redis/redis"
)

type RedisCachingWrapper interface {
	GetCachedValue(k string) string
	SaveCachedValue(k, v string)
}

type redisCachingWrapper struct {
	RedisClient *Redis.Client
}

func NewRedisCachingWrapper() RedisCachingWrapper {
	var redisClient *Redis.Client = ExampleNewClient()
	return &redisCachingWrapper{RedisClient: redisClient}
}

func (rcw *redisCachingWrapper) GetCachedValue(k string) string {
	return ""
}

func (rcw *redisCachingWrapper) SaveCachedValue(k, v string) {

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
