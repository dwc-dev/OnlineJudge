package logic

import (
	"context"

	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionJudgeConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionJudgeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionJudgeConfigLogic {
	return &GetQuestionJudgeConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuestionJudgeConfigLogic) GetQuestionJudgeConfig(in *question.GetQuestionJudgeConfigReq) (*question.GetQuestionJudgeConfigResp, error) {
	q, err := l.svcCtx.QuestionDao.GetQuestionJudgeConfig(in.Id)
	if err != nil {
		return nil, err
	}

	return &question.GetQuestionJudgeConfigResp{
		JudgeConfig: q.JudgeConfig,
	}, nil
}
