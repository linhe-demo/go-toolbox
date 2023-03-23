package gamematching

import "toolbox/common"

// MatchPlayParam 构建匹配者信息
func MatchPlayParam(id int64, rank int64, needNum int64, ch chan int64) *PlayerParam {
	return &PlayerParam{
		Id:             id,
		Rank:           rank,
		Notify:         ch,
		BeginTimestamp: common.CurrentTimestamp(),
		NeedPlayerNum:  needNum,
	}
}

// AddPlayerToPool 向匹配池中新增单个玩家
func (m *MatchPool) AddPlayerToPool(id int64, rank int64, needNum int64, ch chan int64) bool {
	player := MatchPlayParam(id, rank, needNum, ch)
	v, ok := m.PlayMap.LoadOrStore(id, player)
	if v == nil && ok == true {
		m.Size++
	}
	return ok
}

// RemovePlayerOutPool 删除匹配池里的玩家
func (m *MatchPool) RemovePlayerOutPool(id int64) (out bool) {
	if _, out = m.PlayMap.LoadAndDelete(id); out {
		m.Size--
	}
	return out
}

// ReceiveReturnPlayer 将匹配未成功的玩家重新加入匹配池
func (m *MatchPool) ReceiveReturnPlayer(param *PlayerParam) (out bool) {
	if _, out = m.PlayMap.LoadOrStore(param.Id, param); out {
		m.Size++
	}
	return out
}

// NotifyPlayerMatchComplete 通知玩家匹配成功
func (m *PlayerParam) NotifyPlayerMatchComplete(roomId int64) {
	m.Notify <- roomId
}
