package gamematching

import (
	"fmt"
	"toolbox/common"
	"toolbox/pkg/snowflake"
)

// AddPlayerToRoom 向房间里加入玩家信息
func (m *MatchPlayerRoom) AddPlayerToRoom(param *PlayerParam) {
	m.PlayerList = append(m.PlayerList, param)
	m.Size++
}

// DealMatchSuccess 通知房间里的玩家，对局已经找到
func (m *MatchPlayerRoom) DealMatchSuccess() {
	tmpRoomId, _ := snowflake.GetSnowId() //构建房间id
	//向数据库中写入对战房间信息
	//saveInfoToDataBase()
	common.ShowLog(fmt.Sprintf("房间ID为：%d", tmpRoomId))
	for _, v := range m.PlayerList { // 调用用户channel 通知用户匹配完成
		v.NotifyPlayerMatchComplete(tmpRoomId)
	}
}

// InitRoom 初始化对战房间
func InitRoom(roomId int64, list []*PlayerParam, size int64) *MatchPlayerRoom {
	return &MatchPlayerRoom{
		RoomId:     roomId,
		PlayerList: list,
		Size:       size,
	}
}

// RemovePlayerOutRoom 将玩家从房间内移除
func (m *MatchPlayerRoom) RemovePlayerOutRoom(Id int64) {
	j := 0
	for _, v := range m.PlayerList {
		if v.Id != Id {
			m.PlayerList[j] = v
			j++
		} else {
			m.Size--
		}
	}
	m.PlayerList = m.PlayerList[:j]
}
