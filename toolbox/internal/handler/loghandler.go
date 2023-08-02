package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"toolbox/internal/logic"
	"toolbox/internal/svc"
	"toolbox/internal/types"
)

func LogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLogLogic(r.Context(), svcCtx, w, r)
		resp, err := l.Log(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
