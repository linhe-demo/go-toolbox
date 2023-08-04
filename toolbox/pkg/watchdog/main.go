package watchdog

import (
	"context"
	"net/http"
	"toolbox/common"
	"toolbox/internal/models"
	"toolbox/internal/svc"
	"toolbox/tools"
)

func Save(ctx context.Context, svc *svc.ServiceContext, param []LogInfo, w http.ResponseWriter, r *http.Request) {
	for _, v := range param {
		var actionUser = "golang"
		if len(v.ActionUser) != common.Zero {
			actionUser = v.ActionUser
		}
		log := &models.UserLog{
			Ip:         tools.GetOutBoundIpNew(w, r),
			Action:     v.Action,
			ActionUser: actionUser,
		}
		_, _ = svc.UserLogModel.Insert(ctx, log)
	}
}
