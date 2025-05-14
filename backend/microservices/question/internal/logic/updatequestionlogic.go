package logic

import (
	"context"

	"backend/microservices/question/internal/db/model"
	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateQuestionLogic {
	return &UpdateQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateQuestionLogic) UpdateQuestion(in *question.UpdateQuestionReq) (*question.UpdateQuestionResp, error) {
	q := &model.Question{
		ID:           in.QuestionInfo.Id,
		Title:        in.QuestionInfo.Title,
		Content:      in.QuestionInfo.Content,
		Tags:         in.QuestionInfo.Tags,
		Difficulty:   in.QuestionInfo.Difficulty,
		JudgeCase:    in.QuestionInfo.JudgeCase,
		JudgeConfig:  in.QuestionInfo.JudgeConfig,
		VisibleScope: in.QuestionInfo.VisibleScope,
	}

	if err := l.svcCtx.QuestionDao.UpdateQuestion(q); err != nil {
		return nil, err
	}

	return &question.UpdateQuestionResp{}, nil
}
