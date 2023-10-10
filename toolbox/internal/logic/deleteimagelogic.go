package logic

import (
	"context"
	"fmt"
	"os"

	"toolbox/internal/svc"
	"toolbox/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteImageLogic {
	return &DeleteImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteImageLogic) DeleteImage(req *types.DeleteImageRequest) (resp *types.DeleteImageResponse, err error) {
	filePath := fmt.Sprintf("var/www/html%s", req.Name)
	// 检查文件是否存在
	_, err = os.Stat(filePath)
	back := &types.DeleteImageResponse{}
	if err == nil {
		// 文件存在，删除文件
		err := os.Remove(filePath)
		if err != nil {
			back.Path = err.Error()
		}
		back.Path = "ok"
	} else if os.IsNotExist(err) {
		back.Path = "ok"
	} else {
		back.Path = fmt.Sprintf("无法检查文件存在性:%s", err.Error())
	}
	return back, nil
}
