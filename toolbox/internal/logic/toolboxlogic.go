package logic

import (
	"context"
	"toolbox/common"

	"toolbox/internal/svc"
	"toolbox/internal/types"

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
	roomId, _ = <-ch
	if roomId == common.Zero {
		return &types.MatchResponse{}, nil
	}
	resp = &types.MatchResponse{
		RoomId: roomId,
	}
	return resp, nil
}
