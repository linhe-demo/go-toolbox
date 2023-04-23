package redisclient

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"toolbox/internal/config"
)

func Run(config config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%d", config.RedisConf.Host, config.RedisConf.Port)
	opt := &redis.Options{
		Addr:     addr,
		Password: config.RedisConf.Password,
		PoolSize: 10,
		DB:       0,
	}
	return redis.NewClient(opt)
}
