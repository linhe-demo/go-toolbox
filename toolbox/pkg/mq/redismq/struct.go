package redismq

import "github.com/go-redis/redis/v8"

type RedisMessage struct {
	Path string
	Id   int64
}

type RedisListener struct {
	rdb     *redis.Client // Redis client instance
	channel string        // Redis channel name to subscribe to
}
