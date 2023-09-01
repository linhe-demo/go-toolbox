package logic

import (
	"context"
	"net/http"
	"toolbox/common"
	"toolbox/internal/svc"
	"toolbox/internal/types"
	"toolbox/pkg/watchdog"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewLogLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *LogLogic {
	return &LogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *LogLogic) Log(req *types.LogRequest) (resp *types.LogResponse, err error) {
	var param []watchdog.LogInfo
	param = append(param, watchdog.LogInfo{Action: req.Action, ActionUser: req.ActionUser, IP: req.Ip})
	watchdog.Save(l.ctx, l.svcCtx, param, l.w, l.r)
	return &types.LogResponse{
		Result: common.Success,
	}, nil
}
