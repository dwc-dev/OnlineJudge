package user

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"fmt"

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
		fmt.Println("--------------------------------")
		fmt.Println("refresh token error")
		fmt.Println("RefreshToken:", req.RefreshToken)
		fmt.Println("Error:", err)
		fmt.Println("--------------------------------")
		return nil, errors.RefreshTokenInvalid
	}
	fmt.Println("--------------------------------")
	fmt.Println("refresh token success")
	fmt.Println("AccessToken:", rpcResp.AccessToken)
	fmt.Println("RefreshToken:", rpcResp.RefreshToken)
	fmt.Println("--------------------------------")
	return &types.RefreshTokenResp{
		AccessToken:  rpcResp.AccessToken,
		RefreshToken: rpcResp.RefreshToken,
	}, nil
}
