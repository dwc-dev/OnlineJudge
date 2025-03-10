package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
}

type MysqlConfig struct {
	Datasource     string
	ConnectTimeout int
}
