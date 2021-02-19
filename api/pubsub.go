package api

import (
	"github.com/go-redis/redis/v8"
)

type PubSub struct {
}

var Redis *redis.Client

func CreateClient() {
	opt, err := redis.ParseURL("redis://localhost:6364/0")
	if err != nil {
		panic(err)
	}
	redis := redis.NewClient(opt)
	Redis = redis
}
