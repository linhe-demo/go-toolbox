package redismq

import (
	"context"
	"toolbox/internal/config"
	"toolbox/internal/svc"
)

func Consume(c config.Config, context context.Context, ctx *svc.ServiceContext) {
	go func() {
		for true {
			SubscribeMessage(c, context, ctx)
		}
	}()
}
