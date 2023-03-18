package redis

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

var RC *redis.Client

func Init() {
	RC = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RC.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
