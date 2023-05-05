package watchdog

import (
	"context"
	"toolbox/internal/models"
	"toolbox/internal/svc"
	"toolbox/tools"
)

func Save(ctx context.Context, svc *svc.ServiceContext, param []LogInfo) {
	for _, v := range param {
		log := &models.UserLog{
			Ip:         tools.GetOutBoundIP(),
			Action:     v.Action,
			ActionUser: "system",
		}
		_, _ = svc.UserLogModel.Insert(ctx, log)
	}
}
