package svc

import (
	"user/internal/config"
	"user/internal/db"
	"user/internal/db/model/account"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel account.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := db.NewMysqlConn(c.MysqlConfig)
	return &ServiceContext{
		Config:     c,
		UsersModel: account.NewUsersModel(mysqlConn),
	}
}
