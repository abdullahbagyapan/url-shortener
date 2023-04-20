package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Adapter struct {
}

var ctx = context.Background()
var redisClient *redis.Client

func NewAdapter() *Adapter {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "test",
		DB:       0,
	})

	redisClient = rdb

	log.Print("Successfully created redis client")

	return new(Adapter)
}

func (A Adapter) SetToken(token string) error {
	err := redisClient.Set(ctx, token, 1, time.Hour*6).Err()

	if err != redis.Nil {
		return err
	}
	return nil

}

func (A Adapter) IsHasToken(token string) bool {
	hasToken, err := redisClient.Get(ctx, token).Bool()

	if err != nil {
		return false
	}

	return hasToken
}
