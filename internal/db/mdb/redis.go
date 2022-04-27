package mdb

import (
	"api-server/config"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type myRedis struct {
	client *redis.Client
}

func NewRedis(config *config.Redis) (IRedisHandler, error) {
	redisAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: config.Password,
		DB:       config.DB,
	})

	pong, err := client.Ping().Result()
	fmt.Println("Ping redis..... ", pong)
	if err != nil {
		return nil, err
	}
	return &myRedis{
		client: client,
	}, nil
}

func (redis *myRedis) Set(key string, data *[]interface{}, expTime int) error {
	expireTime := time.Duration(expTime)
	err := redis.client.Set(key, data, expireTime).Err()
	return err
}

func (redis *myRedis) Get(key string) (interface{}, error) {
	val, err := redis.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
