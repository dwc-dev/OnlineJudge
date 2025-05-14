package logic

import (
	"context"

	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionListLogic {
	return &GetQuestionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuestionListLogic) GetQuestionList(in *question.GetQuestionListReq) (*question.GetQuestionListResp, error) {
	questions, total, err := l.svcCtx.QuestionDao.GetQuestionList(in.Page, in.PageSize, in.Filter, in.Col)
	if err != nil {
		return nil, err
	}

	res := make([]*question.QuestionInfo, 0)
	for _, q := range questions {
		res = append(res, &question.QuestionInfo{
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
		})
	}

	return &question.GetQuestionListResp{
		QuestionList: res,
		Total:        total,
		Page:         in.Page,
		PageSize:     in.PageSize,
	}, nil
}
