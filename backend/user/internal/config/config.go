package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
	CORSOrigins []string
	JWT         JWT
}

type MysqlConfig struct {
	Datasource     string
	ConnectTimeout int
}

type JWT struct {
	Secret string
	Expire int64
}
