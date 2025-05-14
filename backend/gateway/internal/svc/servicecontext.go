package svc

import (
	"backend/gateway/internal/config"
	middleware2 "backend/gateway/internal/middleware"
	"backend/microservices/competition/competitionclient"
	"backend/microservices/judge/judgeclient"
	"backend/microservices/question/questionclient"
	"backend/microservices/user/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	AccessTokenAuth rest.Middleware
	AdminAuth                rest.Middleware
	CheckContestAccess       rest.Middleware
	RequireContestPermission rest.Middleware
	RequireContestStarted    rest.Middleware
	RequireContestRunning    rest.Middleware
	BlockDuringCompetition   rest.Middleware
	UserRpc                  userclient.User
	QuestionRpc              questionclient.Question
	JudgeRpc                 judgeclient.Judge
	CompetitionRpc           competitionclient.Competition
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                   c,
		AccessTokenAuth:          middleware2.NewAccessTokenAuthMiddleware(c.JWTConf.AccessSecret, userclient.NewUser(zrpc.MustNewClient(c.UserRpc))).Handle,
		AdminAuth:                middleware2.NewAdminAuthMiddleware().Handle,
		CheckContestAccess:       middleware2.NewCheckContestAccessMiddleware(competitionclient.NewCompetition(zrpc.MustNewClient(c.CompetitionRpc))).Handle,
		RequireContestPermission: middleware2.NewRequireContestPermissionMiddleware().Handle,
		RequireContestStarted:    middleware2.NewRequireContestStartedMiddleware().Handle,
		RequireContestRunning:    middleware2.NewRequireContestRunningMiddleware().Handle,
		BlockDuringCompetition:   middleware2.NewBlockDuringCompetitionMiddleware(competitionclient.NewCompetition(zrpc.MustNewClient(c.CompetitionRpc))).Handle,
		UserRpc:                  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		QuestionRpc:              questionclient.NewQuestion(zrpc.MustNewClient(c.QuestionRpc)),
		JudgeRpc:                 judgeclient.NewJudge(zrpc.MustNewClient(c.JudgeRpc)),
		CompetitionRpc:           competitionclient.NewCompetition(zrpc.MustNewClient(c.CompetitionRpc)),
	}
}
