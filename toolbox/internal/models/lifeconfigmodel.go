package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LifeConfigModel = (*customLifeConfigModel)(nil)

type (
	// LifeConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLifeConfigModel.
	LifeConfigModel interface {
		lifeConfigModel
	}

	customLifeConfigModel struct {
		*defaultLifeConfigModel
	}
)

// NewLifeConfigModel returns a model for the database table.
func NewLifeConfigModel(conn sqlx.SqlConn, c cache.CacheConf) LifeConfigModel {
	return &customLifeConfigModel{
		defaultLifeConfigModel: newLifeConfigModel(conn, c),
	}
}
