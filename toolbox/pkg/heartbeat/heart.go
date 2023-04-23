package heartbeat

import (
	"fmt"
	"sync"
	"time"
	"toolbox/common"
	"toolbox/pkg/gamematching"
)

// Init 初始化用户心跳配置
func Init(id int64) *UserHeartStruct {
	return &UserHeartStruct{
		Id:        id,
		Timestamp: time.Now().Unix(),
		Alive:     make(chan int64),
	}
}

// InitHeartPool 初始化用户心跳池
func InitHeartPool() *HeartPool {
	return &HeartPool{
		HeartMap:  sync.Map{},
		Size:      common.Zero,
		Timestamp: time.Now().Unix(),
	}
}

// AddHeatBeat 向用户心跳池里新增用户
func (m *HeartPool) AddHeatBeat(id int64, pool *gamematching.MatchPool) {
	fmt.Println("I am here three times!")
	if v, ok := m.HeartMap.Load(id); ok {
		fmt.Println("I am here once!")
		tmpUser := v.(*UserHeartStruct)
		tmpUser.Alive <- common.Zero
		m.HeartMap.CompareAndSwap(id, v, tmpUser)
	} else {
		fmt.Println("I am here twice!")
		tmpUser := Init(id)
		m.HeartMap.LoadOrStore(id, tmpUser)
		m.Size++
		m.CheckUserIAlive(id, common.WaitSecond, pool)
	}
}

// CheckUserIAlive 检查用户是否还在线
func (m *HeartPool) CheckUserIAlive(id int64, waitSecond int64, pool *gamematching.MatchPool) {
	online := true
	for online {
		if tmpUser, ok := m.HeartMap.Load(id); ok {
			select {
			case <-tmpUser.(*UserHeartStruct).Alive:
				common.ShowLog(fmt.Sprintf("用户 id： %d 还活着", id))
			case <-time.After(time.Duration(waitSecond) * time.Second):
				common.ShowLog(fmt.Sprintf("用户 id： %d 下线", id))
				online = false
				//m.HeartMap.Delete(id)
				//m.Size--
				////移除匹配池中的下线玩家
				//pool.RemovePlayerOutPool(id)
			}
		} else {
			online = false
		}
	}
}
