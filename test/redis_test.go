package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestConnectionRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	s, err := rdb.Get("user").Result()
	if err != nil {
		return
	}
	fmt.Println("user:", s)

	err = rdb.Set("age", "刘诗琪", time.Second*60).Err()
	if err != nil {
		return
	}
}
