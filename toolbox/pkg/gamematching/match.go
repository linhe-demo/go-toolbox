package gamematching

import (
	"fmt"
	"math"
	"toolbox/common"
	"toolbox/pkg/calculate"
)

// MatchProcess 匹配玩家方法
func MatchProcess(pool *MatchPool) {
	var (
		//bTimestamp = common.CurrentTimestamp()
		rankMap = make(map[int64][]*PlayerParam)
	)

	//将玩家按照rank分 分组
	pool.PlayMap.Range(func(k, v any) bool {
		//移除等待超过特定时间的玩家
		tmpUser := v.(*PlayerParam)
		if common.CurrentTimestamp()-tmpUser.BeginTimestamp > common.MaxWaitTime {
			common.ShowLog(fmt.Sprintf("玩家等待时间超过系统最大值：直接移除！玩家id: %v", tmpUser.Id))
			pool.RemovePlayerOutPool(tmpUser.Id)
			tmpUser.NotifyPlayerMatchComplete(common.Zero)
		} else {
			if _, ok := rankMap[tmpUser.Rank]; ok {
				rankMap[tmpUser.Rank] = append(rankMap[tmpUser.Rank], tmpUser)
			} else {
				var tmpPlayerList []*PlayerParam
				tmpPlayerList = append(tmpPlayerList, tmpUser)
				rankMap[tmpUser.Rank] = tmpPlayerList
			}
		}
		return true
	})
	/**
	匹配分组后的玩家
	取同分段等待时间最长的玩家（随着时间推移该玩家的匹配区间是最大的，如果该玩家匹配不到同分段其他玩家也匹配不到
	*/
	for _, v := range rankMap {
		matchBtn := true
		for matchBtn {
			var (
				longestWaitPlayer = InitPlay(common.Zero, common.Zero, common.Zero, make(chan int64))
				min               int64
				max               int64
				rankNum           int64
				needNum           int64
				matchRoom         = InitRoom(common.Zero, []*PlayerParam{}, common.Zero)
			)

			for _, rankPlayer := range v {
				if longestWaitPlayer.Id == common.Zero {
					longestWaitPlayer = rankPlayer
				} else if rankPlayer.BeginTimestamp < longestWaitPlayer.BeginTimestamp {
					longestWaitPlayer = rankPlayer
				}
			}
			if longestWaitPlayer.Id == common.Zero {
				break
			}
			waitTime := float64((common.CurrentTimestamp() - longestWaitPlayer.BeginTimestamp) / 1000)
			common.ShowLog(fmt.Sprintf("开始匹配rank分为：%v 的玩家对局 当前耗时：%v", longestWaitPlayer.Rank, waitTime))

			//随着时间推移慢慢扩大匹配范围
			point := math.Pow(waitTime, common.DefaultSecond)
			point = calculate.AccuracyRound(calculate.AccuracyAdd(common.DefaultWaiteTime, point), common.Zero)
			point = math.Min(point, common.DefaultRankRange)

			tmpRank := longestWaitPlayer.Rank - int64(point)
			if tmpRank >= common.Zero {
				min = tmpRank
			}
			max = longestWaitPlayer.Rank + int64(point)
			common.ShowLog(fmt.Sprintf("本次匹配rank上限分数为：%v 下限分数为：%v", max, min))
			rankNum = longestWaitPlayer.Rank
			needNum = longestWaitPlayer.NeedPlayerNum

			//从此玩家rank分为起点向两边扩大范围
			for rankUp, rankDown := rankNum, rankNum; rankUp <= max || rankDown >= min; rankUp, rankDown = rankUp+common.One, rankDown-common.One {
				var (
					thisPlayers []*PlayerParam
					tmpNum      int64
				)
				if val, ok := rankMap[rankUp]; ok {
					thisPlayers = val
				}
				if rankUp != rankDown && rankDown > common.Zero {
					if val, ok := rankMap[rankDown]; ok {
						thisPlayers = append(thisPlayers, val...)
					}
				}
				tmpNum = int64(len(thisPlayers))
				if tmpNum > common.Zero {
					if matchRoom.Size < needNum {
						for _, target := range thisPlayers {
							//检查用户是否还在线
							if res := target.CheckPlayerOnLine(); !res {
								//将玩家从匹配池中移除
								if r := pool.RemovePlayerOutPool(target.Id); r {
									//将玩家从房间里移除
									matchRoom.RemovePlayerOutRoom(target.Id)
								}
							}

							if target.Id != longestWaitPlayer.Id { //排除玩家自己
								if matchRoom.Size < needNum {
									//将目标玩家从匹配池中移除
									res := pool.RemovePlayerOutPool(target.Id)
									if res == true { //将匹配到的玩家加入房间
										matchRoom.AddPlayerToRoom(target)
									}
									common.ShowLog(fmt.Sprintf("匹配到玩家 玩家id：%v 玩家rank分：%v", target.Id, target.Rank))
								} else {
									break
								}
							}
						}
					}
				}
			}

			// 匹配人数已满
			if matchRoom.Size == needNum {
				//将玩家自己从匹配池中移除
				res := pool.RemovePlayerOutPool(longestWaitPlayer.Id)
				if res == true { //将玩家自己加入房间
					matchRoom.AddPlayerToRoom(longestWaitPlayer)
				}
				common.ShowLog(fmt.Sprintf("对局房间已找到！房间号：%d", matchRoom.RoomId))
				//匹配结束，开始通知等待的玩家
				matchRoom.DealMatchSuccess()
			} else {
				//等待时间最长玩家匹配不到，其他后续玩家不需匹配
				matchBtn = false
				common.ShowLog(fmt.Sprintf("本次匹配失败！等待下一次匹配 等待时间最长玩家id：%d", longestWaitPlayer.Id))

				//将拿出的玩家归还
				for _, v := range matchRoom.PlayerList {
					pool.ReceiveReturnPlayer(v)
				}
			}
		}
	}
	//common.ShowLog(fmt.Sprintf("本次匹配耗时：%v 秒", (common.CurrentTimestamp()-bTimestamp)/1000))
}
