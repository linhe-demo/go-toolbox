package rocketmq

type RocketMq struct {
	Producer string
	Consumer string
}

type Message struct {
	Time string
	Path string
	Id   int64
}
