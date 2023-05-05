package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLogModel = (*customUserLogModel)(nil)

type (
	// UserLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLogModel.
	UserLogModel interface {
		userLogModel
	}

	customUserLogModel struct {
		*defaultUserLogModel
	}
)

// NewUserLogModel returns a model for the database table.
func NewUserLogModel(conn sqlx.SqlConn, c cache.CacheConf) UserLogModel {
	return &customUserLogModel{
		defaultUserLogModel: newUserLogModel(conn, c),
	}
}
