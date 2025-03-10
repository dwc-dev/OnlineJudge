package db

import (
	"context"
	"time"
	"user/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func NewMysqlConn(mysqlConfig config.MysqlConfig) sqlx.SqlConn {
	mysqlConn := sqlx.NewMysql(mysqlConfig.Datasource)
	rawDB, err := mysqlConn.RawDB()
	if err != nil {
		panic(err)
	}
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(mysqlConfig.ConnectTimeout))
	defer cancel() // 防止资源泄漏
	err = rawDB.PingContext(cxt)
	if err != nil {
		panic(err)
	}
	return mysqlConn
}
