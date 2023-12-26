package rocketmq

import (
	"time"
	"toolbox/internal/config"
)

func Run() {
}

func Consume(config config.Config) {
	go func() {
		for true {
			SubscribeMessage(config)
			time.Sleep(time.Hour)
		}
	}()
}
