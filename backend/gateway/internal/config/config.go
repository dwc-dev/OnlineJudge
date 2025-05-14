package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CORSOrigins    []string
	UserRpc        zrpc.RpcClientConf
	QuestionRpc    zrpc.RpcClientConf
	JudgeRpc       zrpc.RpcClientConf
	JWTConf        JWTConf
	CompetitionRpc zrpc.RpcClientConf
	AIServiceURL   string
}

type JWTConf struct {
	AccessSecret string
}
