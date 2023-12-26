package rocketmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"time"
	"toolbox/exception"
	"toolbox/internal/config"
)

const (
	Topic = "LinHe-demo-topic"
	Delay = 3 * time.Second
	Group = "LinHe-demo-group"
)

func SendMessage(config config.Config, message Message) error {
	CreateTopic(config, Topic)
	return SendSyncMessage(config, Topic, message)
}

func CreateTopic(config config.Config, topicName string) {
	endPoint := []string{fmt.Sprintf("%s:%d", config.RocketMqConf.Host, config.RocketMqConf.Port)}
	// 创建主题
	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(endPoint)))
	if err != nil {
		fmt.Printf("connection error: %s\n", err.Error())
	}
	err = testAdmin.CreateTopic(context.Background(), admin.WithTopicCreate(topicName))
	if err != nil {
		fmt.Printf("createTopic error: %s\n", err.Error())
	}
}

func SendSyncMessage(config config.Config, topic string, message Message) error {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{fmt.Sprintf("%s:%d", config.RocketMqConf.Host, config.RocketMqConf.Port)})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		return exception.NewCodeError(exception.RocketMqCode, fmt.Sprintf("start producer error: %s", err.Error()))
	}
	tmp, _ := json.Marshal(message)
	msg := primitive.NewMessage(topic, tmp)
	msg.WithDelayTimeLevel(3)
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		return exception.NewCodeError(exception.RocketMqCode, fmt.Sprintf("send message error: %s\n", err))
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
	err = p.Shutdown()
	if err != nil {
		return exception.NewCodeError(exception.RocketMqCode, fmt.Sprintf("shutdown producer error: %s", err.Error()))
	}
	return nil
}

func SubscribeMessage(config config.Config) {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(Group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{fmt.Sprintf("%s:%d", config.RocketMqConf.Host, config.RocketMqConf.Port)})),
	)
	err := c.Subscribe(Topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Println(msgs[0].Message.String())
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("Shutdown Consumer error: %s", err.Error())
	}
}
