package ai

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CodeCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCodeCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CodeCheckLogic {
	return &CodeCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CodeCheckLogic) CodeCheck(req *types.CodeCheckReq) (resp *types.CodeCheckResp, err error) {
	// todo: add your logic here and delete this line

	return
}
