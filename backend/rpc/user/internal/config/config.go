package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql Mysql
	JWT   JWT
}

type Mysql struct {
	DataSource string
}

type JWT struct {
	Secret string
	Expire int64
}
