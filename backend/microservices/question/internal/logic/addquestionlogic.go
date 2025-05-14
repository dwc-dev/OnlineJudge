package logic

import (
	"context"

	"backend/microservices/question/internal/db/model"
	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddQuestionLogic {
	return &AddQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddQuestionLogic) AddQuestion(in *question.AddQuestionReq) (*question.AddQuestionResp, error) {
	q := &model.Question{
		Title:        in.QuestionInfo.Title,
		Content:      in.QuestionInfo.Content,
		Tags:         in.QuestionInfo.Tags,
		Difficulty:   in.QuestionInfo.Difficulty,
		SubmitNum:    0,
		AcceptedNum:  0,
		JudgeCase:    in.QuestionInfo.JudgeCase,
		JudgeConfig:  in.QuestionInfo.JudgeConfig,
		VisibleScope: in.QuestionInfo.VisibleScope,
	}

	if err := l.svcCtx.QuestionDao.AddQuestion(q); err != nil {
		return nil, err
	}

	return &question.AddQuestionResp{}, nil
}
