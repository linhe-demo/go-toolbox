// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

var (
	userLogFieldNames          = builder.RawFieldNames(&UserLog{})
	userLogRows                = strings.Join(userLogFieldNames, ",")
	userLogRowsExpectAutoSet   = strings.Join(stringx.Remove(userLogFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userLogRowsWithPlaceHolder = strings.Join(stringx.Remove(userLogFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserLogIdPrefix = "cache:userLog:id:"
)

type (
	userLogModel interface {
		Insert(ctx context.Context, data *UserLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserLog, error)
		Update(ctx context.Context, data *UserLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserLogModel struct {
		sqlc.CachedConn
		table string
	}

	UserLog struct {
		Id         int64  `db:"id"`
		Ip         string `db:"ip"`          // 用户ip
		Action     string `db:"action"`      // 用户动作
		ActionUser string `db:"action_user"` // 操作人
		CreateTime string `db:"create_time"` // 创建时间
	}
)

func newUserLogModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserLogModel {
	return &defaultUserLogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_log`",
	}
}

func (m *defaultUserLogModel) Delete(ctx context.Context, id int64) error {
	userLogIdKey := fmt.Sprintf("%s%v", cacheUserLogIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userLogIdKey)
	return err
}

func (m *defaultUserLogModel) FindOne(ctx context.Context, id int64) (*UserLog, error) {
	userLogIdKey := fmt.Sprintf("%s%v", cacheUserLogIdPrefix, id)
	var resp UserLog
	err := m.QueryRowCtx(ctx, &resp, userLogIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userLogRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserLogModel) Insert(ctx context.Context, data *UserLog) (sql.Result, error) {

	query := fmt.Sprintf("insert into %s (ip, action, action_user, create_time) values ('%s', '%s', '%s', '%s')", m.table, data.Ip, data.Action, data.ActionUser, data.CreateTime)
	ret, err := m.ExecNoCacheCtx(ctx, query)
	return ret, err
}

func (m *defaultUserLogModel) Update(ctx context.Context, data *UserLog) error {
	userLogIdKey := fmt.Sprintf("%s%v", cacheUserLogIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userLogRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Ip, data.Action, data.ActionUser, data.Id)
	}, userLogIdKey)
	return err
}

func (m *defaultUserLogModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserLogIdPrefix, primary)
}

func (m *defaultUserLogModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userLogRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserLogModel) tableName() string {
	return m.table
}
