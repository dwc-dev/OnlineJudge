package svc

import (
	"backend/api/gateway/internal/config"
	"backend/api/gateway/internal/middleware"
	"backend/rpc/judge/judgeclient"
	"backend/rpc/question/questionclient"
	"backend/rpc/user/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	JWTAuth     rest.Middleware
	UserRpc     userclient.User
	QuestionRpc questionclient.Question
	JudgeRpc    judgeclient.Judge
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		JWTAuth:     middleware.NewJWTAuthMiddleware(c.JWTConf.AccessSecret).Handle,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		QuestionRpc: questionclient.NewQuestion(zrpc.MustNewClient(c.QuestionRpc)),
		JudgeRpc:    judgeclient.NewJudge(zrpc.MustNewClient(c.JudgeRpc)),
	}
}
