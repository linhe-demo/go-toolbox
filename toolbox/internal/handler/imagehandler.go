package handler

import (
	_ "image/jpeg"
	"mime/multipart"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"toolbox/internal/logic"
	"toolbox/internal/svc"
)

func ImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 读取请求中的图片数据
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Failed to retrieve image", http.StatusBadRequest)
			return
		}

		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		name := r.FormValue("name")
		operateType := r.FormValue("type")

		l := logic.NewImageLogic(r.Context(), svcCtx, w, r)
		resp, err := l.Image(file, header, name, operateType)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
