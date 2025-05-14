package ai

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionSessionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuestionSessionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionSessionsLogic {
	return &GetQuestionSessionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionSessionsLogic) GetQuestionSessions(req *types.GetQuestionSessionsReq) (resp *types.GetQuestionSessionsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
