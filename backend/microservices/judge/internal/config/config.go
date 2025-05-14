package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	QuestionRpc zrpc.RpcClientConf
	Mysql       Mysql
}
type Mysql struct {
	DataSource string
}
