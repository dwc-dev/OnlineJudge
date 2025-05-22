package user

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/common/errors"
	"backend/microservices/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.RefreshTokenResp, err error) {
	rpcResp, err := l.svcCtx.UserRpc.RefreshToken(l.ctx, &userclient.RefreshTokenReq{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return nil, errors.RefreshTokenInvalid
	}
	return &types.RefreshTokenResp{
		AccessToken:  rpcResp.AccessToken,
		RefreshToken: rpcResp.RefreshToken,
	}, nil
}
