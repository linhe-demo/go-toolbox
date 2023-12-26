package mysqlclient

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"toolbox/internal/config"
)

func LogRun(config config.Config) sqlx.SqlConn {
	return sqlx.NewMysql(fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		config.MysqlConf.User, config.MysqlConf.Password, config.MysqlConf.Host, config.MysqlConf.Port, config.MysqlConf.LogDbName, "Asia%2FShanghai"))
}

func LifeRun(config config.Config) sqlx.SqlConn {
	return sqlx.NewMysql(fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		config.MysqlConf.User, config.MysqlConf.Password, config.MysqlConf.Host, config.MysqlConf.Port, config.MysqlConf.LifeDbName, "Asia%2FShanghai"))
}
