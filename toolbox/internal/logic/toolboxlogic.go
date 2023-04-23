package logic

import (
	"context"
	"time"
	"toolbox/common"
	"toolbox/internal/svc"
	"toolbox/internal/types"
	"toolbox/pkg/gamematching"
	"toolbox/pkg/heartbeat"

	"github.com/zeromicro/go-zero/core/logx"
)

type ToolboxLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewToolboxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ToolboxLogic {
	return &ToolboxLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ToolboxLogic) Toolbox(req *types.MatchRequest) (resp *types.MatchResponse, err error) {
	var (
		ch     = make(chan int64, 1)
		roomId int64
	)
	l.svcCtx.Pool.AddPlayerToPool(req.UserId, req.Rank, req.NeedNum, ch)
	l.svcCtx.Heart.AddHeatBeat(req.UserId, l.svcCtx.Pool)
	//开启自主上报
	ReportUserAlive(req.UserId, l.svcCtx.Heart, l.svcCtx.Pool)
	roomId, _ = <-ch
	// 关闭通道
	defer close(ch)
	if roomId == common.Zero {
		return &types.MatchResponse{}, nil
	}
	resp = &types.MatchResponse{
		RoomId: roomId,
	}
	return resp, nil
}

func ReportUserAlive(id int64, heart *heartbeat.HeartPool, pool *gamematching.MatchPool) {
	println("here")
	myTimer := time.NewTimer(time.Second * 2)
	for {
		select {
		case <-myTimer.C:
			heart.AddHeatBeat(id, pool)
			myTimer.Reset(time.Second * 2) // 每次使用完后需要人为重置下
		}
	}
}
