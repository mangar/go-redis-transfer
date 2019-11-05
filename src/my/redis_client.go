package my

import (
	"github.com/go-redis/redis"
)

// RedisClient
//
//
func RedisClient(host string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
