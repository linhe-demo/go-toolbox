package logic

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"toolbox/common"
	"toolbox/internal/svc"
	"toolbox/internal/types"
	"toolbox/pkg/image"
	"toolbox/pkg/watchdog"
	"toolbox/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewImageLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *ImageLogic {
	return &ImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *ImageLogic) Image(file multipart.File, header *multipart.FileHeader, name string) (resp *types.CompressionResponse, err error) {
	var param []watchdog.LogInfo
	//保存图片
	path := tools.SaveFile(file, header)
	//压缩图片
	image.CompressionImage(path, 0.5, name)
	// 写入日志
	param = append(param, watchdog.LogInfo{Action: "compressPhoto"})
	watchdog.Save(l.ctx, l.svcCtx, param, l.w, l.r)
	//返回图片下载地址
	resp = &types.CompressionResponse{
		Path: fmt.Sprintf("%s%s.jpg", common.DownloadFilePath, name),
	}
	return resp, nil
}
