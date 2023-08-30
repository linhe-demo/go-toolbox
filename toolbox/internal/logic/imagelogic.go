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

func (l *ImageLogic) Image(file multipart.File, header *multipart.FileHeader, name string, operateType string) (resp *types.CompressionResponse, err error) {
	var (
		param   []watchdog.LogInfo
		newPath string
	)
	//保存图片
	path := tools.SaveFile(file, header)
	switch operateType {
	case "":
		//压缩图片
		image.CompressionImage(path, 0.5, name)
		newPath = fmt.Sprintf("%s%s.jpg", common.DownloadFilePath, name)
		param = append(param, watchdog.LogInfo{Action: "compressPhoto"})
	case "1":
		//图片转pdf
		image.TransferToPdf(path, name)
		newPath = fmt.Sprintf("%s%s.pdf", common.DownloadFilePath, name)
		param = append(param, watchdog.LogInfo{Action: "imgToPdf"})
	default:
		newPath = "咱不支持该类型转换"
	}
	// 写入日志
	watchdog.Save(l.ctx, l.svcCtx, param, l.w, l.r)
	//返回图片下载地址
	resp = &types.CompressionResponse{
		Path: newPath,
	}
	return resp, nil
}
