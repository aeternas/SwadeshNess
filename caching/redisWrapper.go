package caching

import (
	"errors"
	"fmt"
	Configuration "github.com/aeternas/SwadeshNess/configuration"
	Redis "github.com/go-redis/redis"
)

type RedisCachingWrapper interface {
	GetCachedValue(k string) (string, error)
	SaveCachedValue(k, v string) error
}

type redisCachingWrapper struct {
	RedisClient *Redis.Client
}

func NewRedisCachingWrapper(c *Configuration.Configuration) RedisCachingWrapper {
	var redisClient *Redis.Client = NewClient(c)
	return &redisCachingWrapper{RedisClient: redisClient}
}

func (rcw *redisCachingWrapper) GetCachedValue(k string) (string, error) {
	val, err := rcw.RedisClient.Get(k).Result()
	if err != nil {
		return "", err
	}

	if err == redis.Nil {
		errorMessage := fmt.Sprintf("Key %v doesn't exist", err)
		return "", errorMessage
	}
	return val, nil
}

func (rcw *redisCachingWrapper) SaveCachedValue(k, v string) error {
	err := rcw.RedisClient.Set(k, v, 0).Err()
	return err
}

func NewClient(c *Configuration.Configuration) *Redis.Client {
	address := fmt.Sprintf("%s:6379", c.EEndpoints.RedisAddress)
	client := Redis.NewClient(&Redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
