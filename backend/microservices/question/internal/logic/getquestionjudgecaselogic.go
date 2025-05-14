package logic

import (
	"context"

	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionJudgeCaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionJudgeCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionJudgeCaseLogic {
	return &GetQuestionJudgeCaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuestionJudgeCaseLogic) GetQuestionJudgeCase(in *question.GetQuestionJudgeCaseReq) (*question.GetQuestionJudgeCaseResp, error) {
	q, err := l.svcCtx.QuestionDao.GetQuestionJudgeCase(in.Id)
	if err != nil {
		return nil, err
	}

	return &question.GetQuestionJudgeCaseResp{
		JudgeCase: q.JudgeCase,
	}, nil
}
