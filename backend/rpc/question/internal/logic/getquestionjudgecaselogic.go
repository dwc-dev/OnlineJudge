package logic

import (
	"context"

	"backend/rpc/question/internal/model"
	"backend/rpc/question/internal/svc"
	"backend/rpc/question/pb/question"

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

func (l *GetQuestionJudgeCaseLogic) GetQuestionJudgeCase(in *question.QuestionRequest) (*question.QuestionJudgeCaseResponse, error) {
	var q model.Question
	result := l.svcCtx.DB.Where("id = ?", in.Id).First(&q)
	if result.Error != nil {
		return nil, result.Error
	}
	return &question.QuestionJudgeCaseResponse{
		JudgeCase: q.JudgeCase,
	}, nil
}
