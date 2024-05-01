package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var memory *Memo

func Conn() *Memo {
	return memory
}

func init() {
	memory = &Memo{redisPool()}
}

func redisPool() *redis.Client {
	rdsAddr := func() string {
		addr := os.Getenv("REDIS_ADDR")
		if addr == "" {
			addr = "localhost:6379"
		}

		return addr
	}

	rds := redis.NewClient(&redis.Options{
		Addr: rdsAddr(),
	})

	_, err := rds.
		Ping(context.Background()).
		Result()

	if err != nil {
		panic(err)
	}

	log.Println("Connected to redis")

	return rds
}
