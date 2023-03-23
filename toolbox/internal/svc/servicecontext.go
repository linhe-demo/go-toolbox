package svc

import (
	"toolbox/internal/config"
	"toolbox/pkg/gamematching"
)

type ServiceContext struct {
	Config config.Config
	Pool   *gamematching.MatchPool
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Pool:   gamematching.Run(),
	}
}
