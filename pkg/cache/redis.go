package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"sync"
)

var (
	once    sync.Once
	rdsPool *redis.Client
)

func rdsClient() *redis.Client {
	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr: rdsAddr(),
		})

		cmd := rdb.Ping(context.Background())
		if _, err := cmd.Result(); err != nil {
			panic(err)
		}

		rdsPool = rdb
	})

	return rdsPool
}
func rdsAddr() string {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	return addr
}
