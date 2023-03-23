package gamematching

import "sync"

// PlayerParam 玩家信息
type PlayerParam struct {
	Id             int64      //玩家id
	Rank           int64      //玩家rank分
	Notify         chan int64 //用户消息通道
	BeginTimestamp int64      //玩家开始时间
	NeedPlayerNum  int64      //需要匹配的玩家数量
}

// MatchPool 匹配池
type MatchPool struct {
	PlayMap sync.Map //玩家池map
	Size    int64    //匹配池玩家人数
}

// MatchPlayerRoom 玩家对战房间
type MatchPlayerRoom struct {
	RoomId     int64          //房间ID
	PlayerList []*PlayerParam // 房间玩家成员切片
	Size       int64          //房间成员数量
}
