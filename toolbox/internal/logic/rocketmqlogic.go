package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"toolbox/common"
	"toolbox/internal/svc"
	"toolbox/internal/types"
	"toolbox/pkg/mq/rocketmq"
)

type RocketMqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRocketMqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RocketMqLogic {
	return &RocketMqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RocketMqLogic) RocketMq(req *types.RocketMqRequest) (resp *types.RocketMqResponse, err error) {
	err = rocketmq.SendMessage(l.svcCtx.Config, rocketmq.Message{
		Time: time.Now().Format(common.TimeFormat),
		Path: req.Path,
		Id:   req.Id,
	})
	if err != nil {
		return &types.RocketMqResponse{
			Result: err.Error(),
		}, nil
	}
	return &types.RocketMqResponse{
		Result: "send success",
	}, nil
}
