package svc

import (
	"github.com/go-redis/redis/v8"
	"toolbox/internal/config"
	"toolbox/internal/models"
	"toolbox/pkg/gamematching"
	"toolbox/pkg/heartbeat"
	"toolbox/pkg/mysqlclient"
	"toolbox/pkg/redisclient"
	"toolbox/pkg/websocket"
)

type ServiceContext struct {
	Config          config.Config
	Pool            *gamematching.MatchPool
	Hub             *websocket.Hub
	Heart           *heartbeat.HeartPool
	RedisClient     *redis.Client
	UserLogModel    models.UserLogModel
	LifeConfigModel models.LifeConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlClient := mysqlclient.Run(c)
	return &ServiceContext{
		Config:          c,
		Pool:            gamematching.Run(),
		Hub:             websocket.Run(),
		Heart:           heartbeat.Run(),
		RedisClient:     redisclient.Run(c),
		UserLogModel:    models.NewUserLogModel(MysqlClient, c.CacheRedis),
		LifeConfigModel: models.NewLifeConfigModel(MysqlClient, c.CacheRedis),
	}
}
