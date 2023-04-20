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
var rdsClient *redis.Client

func NewAdapter() *Adapter {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rdsClient = rdb

	log.Printf("Succesfully connected to the Redis")

	return new(Adapter)

}

func (A Adapter) GetURL(shortURL string) string {
	val := rdsClient.HGet(ctx, shortURL, "long_url").Val()

	return val

}

func (A Adapter) SaveURL(shortURL, longURL string) error {
	err := rdsClient.HSet(ctx, shortURL, "long_url", longURL).Err()

	if err != nil {
		return err
	}

	rdsClient.Expire(ctx, shortURL, time.Minute*20)

	return nil
}
