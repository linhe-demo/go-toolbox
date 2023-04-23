package svc

import (
	"github.com/go-redis/redis/v8"
	"toolbox/internal/config"
	"toolbox/pkg/gamematching"
	"toolbox/pkg/heartbeat"
	"toolbox/pkg/redisclient"
	"toolbox/pkg/websocket"
)

type ServiceContext struct {
	Config      config.Config
	Pool        *gamematching.MatchPool
	Hub         *websocket.Hub
	Heart       *heartbeat.HeartPool
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Pool:        gamematching.Run(),
		Hub:         websocket.Run(),
		Heart:       heartbeat.Run(),
		RedisClient: redisclient.Run(c),
	}
}
