package handler

import (
	"net/http"
	"toolbox/exception"

	"github.com/zeromicro/go-zero/rest/httpx"
	"toolbox/internal/logic"
	"toolbox/internal/svc"
	"toolbox/internal/types"
)

func ToolboxHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MatchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, exception.NewCodeError(exception.ParamCode, err.Error()))
			return
		}

		l := logic.NewToolboxLogic(r.Context(), svcCtx)
		resp, err := l.Toolbox(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
