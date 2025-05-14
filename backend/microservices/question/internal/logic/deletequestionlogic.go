package logic

import (
	"context"

	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteQuestionLogic {
	return &DeleteQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteQuestionLogic) DeleteQuestion(in *question.DeleteQuestionReq) (*question.DeleteQuestionResp, error) {
	if err := l.svcCtx.QuestionDao.DeleteQuestion(in.Id); err != nil {
		return nil, err
	}

	return &question.DeleteQuestionResp{}, nil
}
