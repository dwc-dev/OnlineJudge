package question

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"encoding/json"

	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateQuestionLogic {
	return &UpdateQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateQuestionLogic) UpdateQuestion(req *types.UpdateQuestionReq) (resp *types.UpdateQuestionResp, err error) {
	tagsBytes, err := json.Marshal(req.QuestionInfo.Tags)
	if err != nil {
		return nil, err
	}
	judgeCaseBytes, err := json.Marshal(req.QuestionInfo.JudgeCase)
	if err != nil {
		return nil, err
	}
	judgeConfigBytes, err := json.Marshal(req.QuestionInfo.JudgeConfig)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.QuestionRpc.UpdateQuestion(l.ctx, &question.UpdateQuestionReq{
		QuestionInfo: &question.QuestionInfo{
			Id:           req.QuestionInfo.Id,
			Title:        req.QuestionInfo.Title,
			Content:      req.QuestionInfo.Content,
			Tags:         string(tagsBytes),
			Difficulty:   req.QuestionInfo.Difficulty,
			JudgeCase:    string(judgeCaseBytes),
			JudgeConfig:  string(judgeConfigBytes),
			VisibleScope: req.QuestionInfo.VisibleScope,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateQuestionResp{}, nil
}
