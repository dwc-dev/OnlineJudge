package logic

import (
	"context"

	"backend/microservices/user/internal/svc"
	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.LogoutResp, error) {
	err := l.svcCtx.RedisClient.DeleteAccessTokenJTI(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.DeleteRefreshTokenJTI(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &user.LogoutResp{}, nil
}
