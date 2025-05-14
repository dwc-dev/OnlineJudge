package ai

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSessionMessageHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSessionMessageHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSessionMessageHistoryLogic {
	return &GetSessionMessageHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSessionMessageHistoryLogic) GetSessionMessageHistory(req *types.GetSessionMessageHistoryReq) (resp *types.GetSessionMessageHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
