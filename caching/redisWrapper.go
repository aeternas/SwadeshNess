package caching

import (
	"fmt"
	Configuration "github.com/aeternas/SwadeshNess/configuration"
	redis "github.com/go-redis/redis"
)

type RedisCachingWrapper interface {
	GetCachedValue(k string) (string, error)
	SaveCachedValue(k, v string) error
}

type redisCachingWrapper struct {
	RedisClient *redis.Client
}

func NewRedisCachingWrapper(c *Configuration.Configuration) RedisCachingWrapper {
	var redisClient *redis.Client = NewClient(c)
	return &redisCachingWrapper{RedisClient: redisClient}
}

func (rcw *redisCachingWrapper) GetCachedValue(k string) (string, error) {
	val, err := rcw.RedisClient.Get(k).Result()
	if err != nil {
		return "", fmt.Errorf("Internal Redis Error: ", err)
	}

	if err == redis.Nil {
		errorMessage := fmt.Errorf("Key %v doesn't exist", err)
		return "", errorMessage
	}
	return val, nil
}

func (rcw *redisCachingWrapper) SaveCachedValue(k, v string) error {
	err := rcw.RedisClient.Set(k, v, 0).Err()
	return err
}

func NewClient(c *Configuration.Configuration) *redis.Client {
	address := fmt.Sprintf("%s:6379", c.EEndpoints.RedisAddress)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
