package question

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"encoding/json"

	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionBasicInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuestionBasicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionBasicInfoLogic {
	return &GetQuestionBasicInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionBasicInfoLogic) GetQuestionBasicInfo(req *types.GetQuestionBasicInfoReq) (resp *types.GetQuestionBasicInfoResp, err error) {
	r, err := l.svcCtx.QuestionRpc.GetQuestionInfo(l.ctx, &question.GetQuestionInfoReq{
		Id:     req.Id,
		Col:    []string{"title", "content", "judge_config"},
		Filter: map[string]string{"visible_scope": "public"}, //确保只能查询公开的题目
	})
	if err != nil {
		return nil, err
	}
	var judgeConfig *types.JudgeConfig
	err = json.Unmarshal([]byte(r.QuestionInfo.JudgeConfig), &judgeConfig)
	if err != nil {
		return nil, err
	}
	return &types.GetQuestionBasicInfoResp{
		QuestionInfo: types.QuestionInfo{
			Id:          r.QuestionInfo.Id,
			Title:       r.QuestionInfo.Title,
			Content:     r.QuestionInfo.Content,
			JudgeConfig: judgeConfig,
		},
	}, nil
}
