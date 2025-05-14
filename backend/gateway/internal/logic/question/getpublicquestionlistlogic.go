package question

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"encoding/json"

	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicQuestionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicQuestionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicQuestionListLogic {
	return &GetPublicQuestionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicQuestionListLogic) GetPublicQuestionList(req *types.GetPublicQuestionListReq) (resp *types.GetPublicQuestionListResp, err error) {
	// 确保只能查询公开的题目
	req.Filter["visible_scope"] = "public"

	r, err := l.svcCtx.QuestionRpc.GetQuestionList(l.ctx, &question.GetQuestionListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Filter:   req.Filter,
		Col:      []string{"id", "title", "tags", "difficulty", "submit_num", "accepted_num"},
	})
	if err != nil {
		return nil, err
	}

	res := make([]*types.QuestionInfo, 0)
	for _, q := range r.QuestionList {
		var tags []string
		err = json.Unmarshal([]byte(q.Tags), &tags)
		if err != nil {
			return nil, err
		}
		res = append(res, &types.QuestionInfo{
			Id:          q.Id,
			Title:       q.Title,
			Tags:        tags,
			Difficulty:  q.Difficulty,
			SubmitNum:   q.SubmitNum,
			AcceptedNum: q.AcceptedNum,
		})
	}
	return &types.GetPublicQuestionListResp{
		QuestionInfo: res,
		Total:        r.Total,
		Page:         r.Page,
		PageSize:     r.PageSize,
	}, nil
}
