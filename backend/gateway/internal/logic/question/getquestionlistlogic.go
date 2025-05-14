package question

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"encoding/json"
	"fmt"

	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuestionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionListLogic {
	return &GetQuestionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionListLogic) GetQuestionList(req *types.GetQuestionListReq) (resp *types.GetQuestionListResp, err error) {
	r, err := l.svcCtx.QuestionRpc.GetQuestionList(l.ctx, &question.GetQuestionListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Filter:   req.Filter,
		Col:      nil,
	})
	if err != nil {
		fmt.Println("GetQuestionList error: ", err)
		return nil, err
	}
	res := make([]*types.QuestionInfo, 0)
	for _, q := range r.QuestionList {
		var tags []string
		err = json.Unmarshal([]byte(q.Tags), &tags)
		if err != nil {
			fmt.Println("Unmarshal Tags error: ", err)
			tags = []string{}
		}
		var judgeCase []*types.JudgeCase
		err = json.Unmarshal([]byte(q.JudgeCase), &judgeCase)
		if err != nil {
			fmt.Println("Unmarshal JudgeCase error: ", err)
			judgeCase = []*types.JudgeCase{}
		}
		var judgeConfig *types.JudgeConfig
		err = json.Unmarshal([]byte(q.JudgeConfig), &judgeConfig)
		if err != nil {
			fmt.Println("Unmarshal JudgeConfig error: ", err)
			judgeConfig = &types.JudgeConfig{}
		}
		res = append(res, &types.QuestionInfo{
			Id:           q.Id,
			Title:        q.Title,
			Content:      q.Content,
			Tags:         tags,
			Difficulty:   q.Difficulty,
			SubmitNum:    q.SubmitNum,
			AcceptedNum:  q.AcceptedNum,
			JudgeCase:    judgeCase,
			JudgeConfig:  judgeConfig,
			VisibleScope: q.VisibleScope,
			CreateAt:     q.CreateAt,
			UpdateAt:     q.UpdateAt,
		})
	}
	return &types.GetQuestionListResp{
		QuestionInfo: res,
		Total:        r.Total,
		Page:         r.Page,
		PageSize:     r.PageSize,
	}, nil
}
