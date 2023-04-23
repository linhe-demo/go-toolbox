package heartbeat

import "sync"

type UserHeartStruct struct {
	Id        int64
	Timestamp int64
	Alive     chan int64
}

type HeartPool struct {
	HeartMap  sync.Map
	Size      int64
	Timestamp int64
}
