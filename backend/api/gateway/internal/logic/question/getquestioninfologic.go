package question

import (
	"context"

	"backend/api/gateway/internal/svc"
	"backend/api/gateway/internal/types"
	"backend/rpc/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuestionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionInfoLogic {
	return &GetQuestionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionInfoLogic) GetQuestionInfo(req *types.QuestionReq) (resp *types.QuestionInfoResponse, err error) {
	r, err := l.svcCtx.QuestionRpc.GetQuestionInfo(l.ctx, &question.QuestionRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.QuestionInfoResponse{
		Title:   r.Title,
		Content: r.Content,
	}, nil
}
