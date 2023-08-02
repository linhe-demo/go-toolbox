package logic

import (
	"context"
	"net/http"
	"toolbox/internal/svc"
	"toolbox/internal/types"
	"toolbox/pkg/oauth"
	"toolbox/pkg/watchdog"

	"github.com/zeromicro/go-zero/core/logx"
)

type OcrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewOcrLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *OcrLogic {
	return &OcrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *OcrLogic) Ocr(req *types.OcrRequest) (resp *types.OcrResponse, err error) {
	var param []watchdog.LogInfo
	res, err := oauth.AnalysisPictureText(l.ctx, l.svcCtx, req.Type, req.File, req.FileType, l.w, l.r)
	if err != nil {
		return resp, err
	}
	// 处理数据
	var resultList []string
	for _, v := range res.WordsResult {
		resultList = append(resultList, v.Words)
	}
	// 写入日志
	param = append(param, watchdog.LogInfo{Action: "ocr"})
	watchdog.Save(l.ctx, l.svcCtx, param, l.w, l.r)
	return &types.OcrResponse{
		Result: resultList,
	}, nil
}
