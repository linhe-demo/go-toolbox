package handler

import (
	"fmt"
	"net/http"
	"toolbox/exception"

	"github.com/zeromicro/go-zero/rest/httpx"
	"toolbox/internal/logic"
	"toolbox/internal/svc"
	"toolbox/internal/types"
)

func RocketMqHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RocketMqRequest
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err.Error())
			httpx.Error(w, exception.NewError(exception.ParamCode, err.Error()))
			return
		}

		l := logic.NewRocketMqLogic(r.Context(), svcCtx)
		resp, err := l.RocketMq(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
