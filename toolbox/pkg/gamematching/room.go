package gamematching

// AddPlayerToRoom 向房间里加入玩家信息
func (m *MatchPlayerRoom) AddPlayerToRoom(param *PlayerParam) {
	m.PlayerList = append(m.PlayerList, param)
	m.Size++
}

// DealMatchSuccess 通知房间里的玩家，对局已经找到
func (m *MatchPlayerRoom) DealMatchSuccess() {
	//向数据库中写入对战房间信息
	//saveInfoToDataBase()
	for _, v := range m.PlayerList { // 调用用户channel 通知用户匹配完成
		v.NotifyPlayerMatchComplete(m.RoomId)
	}
}
