package redismq

import (
	"context"
	"toolbox/internal/config"
	"toolbox/internal/svc"
	"toolbox/pkg/image"
)

const (
	RedisQueueKey = "LinHe-demo-queue"
)

func SubscribeMessage(c config.Config, context context.Context, ctx *svc.ServiceContext) {
	pubsub := ctx.RedisClient.Subscribe(context, RedisQueueKey)
	defer pubsub.Close()
	for {
		msg, err := pubsub.ReceiveMessage(context)
		if err != nil {
			return
		}
		image.DealImageFile(c, context, ctx, msg.Payload)
	}
}
