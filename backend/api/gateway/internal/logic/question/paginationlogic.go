package question

import (
	"context"

	"backend/api/gateway/internal/svc"
	"backend/api/gateway/internal/types"
	"backend/rpc/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaginationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginationLogic {
	return &PaginationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaginationLogic) Pagination(req *types.PaginationReq) (resp *types.PaginationResp, err error) {
	r, err := l.svcCtx.QuestionRpc.Pagination(l.ctx, &question.PaginationRequest{Page: req.Page,
		PageSize: req.PageSize,
		Filter:   req.Filter,
	})
	if err != nil {
		return nil, err
	}
	questionList := make([]*types.QuestionBasicInfo, 0)
	for _, question := range r.Questions {
		questionList = append(questionList, &types.QuestionBasicInfo{
			Id:         question.Id,
			Title:      question.Title,
			Tags:       question.Tags,
			AcRate:     question.AcRate,
			Difficulty: question.Difficulty,
		})
	}
	resp = &types.PaginationResp{
		Total:     r.Total,
		Page:      r.Page,
		PageSize:  r.PageSize,
		Questions: questionList,
	}
	return resp, nil
}
