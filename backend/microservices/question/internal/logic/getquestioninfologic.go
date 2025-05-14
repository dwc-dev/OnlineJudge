package logic

import (
	"context"

	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

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

func (l *GetQuestionInfoLogic) GetQuestionInfo(in *question.GetQuestionInfoReq) (*question.GetQuestionInfoResp, error) {
	q, err := l.svcCtx.QuestionDao.GetQuestionInfo(in.Id, in.Filter, in.Col)
	if err != nil {
		return nil, err
	}

	return &question.GetQuestionInfoResp{
		QuestionInfo: &question.QuestionInfo{
			Id:           q.ID,
			Title:        q.Title,
			Content:      q.Content,
			Tags:         q.Tags,
			Difficulty:   q.Difficulty,
			SubmitNum:    q.SubmitNum,
			AcceptedNum:  q.AcceptedNum,
			JudgeCase:    q.JudgeCase,
			JudgeConfig:  q.JudgeConfig,
			VisibleScope: q.VisibleScope,
			CreateAt:     q.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:     q.UpdateAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
