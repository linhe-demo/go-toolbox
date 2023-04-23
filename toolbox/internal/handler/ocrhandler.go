package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"toolbox/internal/logic"
	"toolbox/internal/svc"
	"toolbox/internal/types"
)

func OcrHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OcrRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOcrLogic(r.Context(), svcCtx)
		resp, err := l.Ocr(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
