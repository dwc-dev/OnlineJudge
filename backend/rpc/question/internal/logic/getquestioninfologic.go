package logic

import (
	"context"

	"backend/rpc/question/internal/model"
	"backend/rpc/question/internal/svc"
	"backend/rpc/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionInfoLogic {
	return &GetQuestionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuestionInfoLogic) GetQuestionInfo(in *question.QuestionRequest) (*question.QuestionInfoResponse, error) {
	var q model.Question
	result := l.svcCtx.DB.Where("id = ?", in.Id).First(&q)
	if result.Error != nil {
		return nil, result.Error
	}
	return &question.QuestionInfoResponse{
		Title:   q.Title,
		Content: q.Content,
	}, nil
}
