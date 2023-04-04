package redis

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var RC *redis.Client

const expireTime = 30 * time.Second

func Init() {
	RC = redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: "",
		DB:       0,
	})

	_, err := RC.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func AddData(key string, data string) error {
	_, err := RC.SetNX(context.Background(), key, data, expireTime).Result()
	return err
}

func GetData(key string) ([]byte, error) {
	val, err := RC.Get(context.Background(), key).Bytes()
	return val, err
}
