package gamematching

import (
	"fmt"
	"sync"
	"time"
	"toolbox/common"
)

// Pool 初始化匹配池
var Pool = InitPlayerPoolInfo()

func Run() *MatchPool {

	// 开起协程
	go func() {
		common.ShowLog(fmt.Sprintf("匹配池启动成功，时间：%v", time.Now().Format(common.TimeFormat)))
		// 启用定时器
		myTimer := time.NewTimer(time.Second * 1)
		for {
			select {
			case <-myTimer.C:
				//匹配函数
				MatchProcess(Pool)
				myTimer.Reset(time.Second * 1)
			}
		}
	}()
	return Pool
}

// InitPlayerPoolInfo 匹配池用户信息初始化
func InitPlayerPoolInfo() *MatchPool {
	return &MatchPool{
		PlayMap: sync.Map{},
		Size:    common.Zero,
	}
}
