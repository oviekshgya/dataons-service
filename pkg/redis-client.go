package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

type GenerateAccess struct {
	Time time.Time `json:"Time"`
}

type LoggerRedis struct {
	Code         string      `json:"code"`
	Timestamp    time.Time   `json:"timestamp"`
	Id           int         `json:"id"`
	Repositories string      `json:"repositories"`
	Column       int         `json:"column"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

// GetClient get the redis client
func InitializeRedis() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	if err := c.Ping().Err(); err != nil {
		fmt.Println("Unable to connect to redis " + err.Error())
	}
	client.c = c
	return client
}

func (client *redisClient) GetKey(key string, src interface{}) error {
	val, err := client.c.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}

// SetKey set key
func (client *redisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
