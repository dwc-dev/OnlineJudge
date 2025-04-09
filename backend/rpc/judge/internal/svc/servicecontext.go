package svc

import (
	"backend/rpc/judge/internal/config"
	"backend/rpc/question/questionclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	QuestionRpc questionclient.Question
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		QuestionRpc: questionclient.NewQuestion(zrpc.MustNewClient(c.QuestionRpc)),
	}
}
