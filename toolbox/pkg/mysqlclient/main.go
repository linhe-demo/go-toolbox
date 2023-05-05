package mysqlclient

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"toolbox/internal/config"
)

func Run(config config.Config) sqlx.SqlConn {
	return sqlx.NewMysql(fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		config.MysqlConf.User, config.MysqlConf.Password, config.MysqlConf.Host, config.MysqlConf.Port, config.MysqlConf.DbName, "Asia%2FShanghai"))
}
