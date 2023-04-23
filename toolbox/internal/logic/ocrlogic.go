package logic

import (
	"context"
	"toolbox/internal/svc"
	"toolbox/internal/types"
	"toolbox/pkg/oauth"

	"github.com/zeromicro/go-zero/core/logx"
)

type OcrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOcrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OcrLogic {
	return &OcrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OcrLogic) Ocr(req *types.OcrRequest) (resp *types.OcrResponse, err error) {
	res, err := oauth.AnalysisPictureText(l.ctx, l.svcCtx.RedisClient, l.svcCtx.Config, req.Type, req.File, req.FileType)
	if err != nil {
		return resp, err
	}
	// 处理数据
	var resultList []string
	for _, v := range res.WordsResult {
		resultList = append(resultList, v.Words)
	}
	return &types.OcrResponse{
		Result: resultList,
	}, nil
}
