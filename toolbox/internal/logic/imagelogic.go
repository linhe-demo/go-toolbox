package logic

import (
	"context"
	"mime/multipart"
	"net/http"
	"toolbox/internal/svc"
	"toolbox/internal/types"
	"toolbox/pkg/image"
	"toolbox/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageLogic {
	return &ImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageLogic) Image(file multipart.File, header *multipart.FileHeader, w http.ResponseWriter, name string) (resp *types.CompressionResponse, err error) {
	//保存图片
	path := tools.SaveFile(file, header)
	//压缩图片
	image.CompressionImage(path, 0.5, name)
	//defer func(name string) {
	//	err := os.Remove(name)
	//	if err != nil {
	//
	//	}
	//}(newPath)
	//发送文件给前端
	//tools.DownloadFile(newPath, w, header.Filename)
	resp = &types.CompressionResponse{
		Path: "http://150.158.82.218/images/" + name + ".jpg",
	}
	return resp, nil
}
