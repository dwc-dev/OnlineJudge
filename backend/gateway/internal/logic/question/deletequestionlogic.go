package question

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteQuestionLogic {
	return &DeleteQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteQuestionLogic) DeleteQuestion(req *types.DeleteQuestionReq) (resp *types.DeleteQuestionResp, err error) {
	_, err = l.svcCtx.QuestionRpc.DeleteQuestion(l.ctx, &question.DeleteQuestionReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteQuestionResp{}, nil
}
