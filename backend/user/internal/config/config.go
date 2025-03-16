package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
	CORSOrigins []string
}

type MysqlConfig struct {
	Datasource     string
	ConnectTimeout int
}
