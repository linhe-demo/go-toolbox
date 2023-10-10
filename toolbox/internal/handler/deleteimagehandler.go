package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"toolbox/internal/logic"
	"toolbox/internal/svc"
	"toolbox/internal/types"
)

func DeleteImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteImageLogic(r.Context(), svcCtx)
		resp, err := l.DeleteImage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
