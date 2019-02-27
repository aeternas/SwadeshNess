package caching

import (
	"fmt"
	Configuration "github.com/aeternas/SwadeshNess/configuration"
	redis "github.com/go-redis/redis"
)

const (
	REDIS_PORT = "6379"
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
	if err == redis.Nil {
		errorMessage := fmt.Errorf("Error extracting key %s", err)
		return "", errorMessage
	}
	if err != nil {
		return "", fmt.Errorf("Internal Redis Error: %v", err)
	}
	return val, nil
}

func (rcw *redisCachingWrapper) SaveCachedValue(k, v string) error {
	err := rcw.RedisClient.Set(k, v, 0).Err()
	return err
}

func NewClient(c *Configuration.Configuration) *redis.Client {
	address := fmt.Sprintf("%s:%s", c.EEndpoints.RedisAddress, REDIS_PORT)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
