package logic

import (
	"context"

	"backend/microservices/user/internal/svc"
	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessTokenJTICheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccessTokenJTICheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessTokenJTICheckLogic {
	return &AccessTokenJTICheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccessTokenJTICheckLogic) AccessTokenJTICheck(in *user.AccessTokenJTICheckReq) (*user.AccessTokenJTICheckResp, error) {
	storedAccessTokenJTI, err := l.svcCtx.RedisClient.GetAccessTokenJTI(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if storedAccessTokenJTI != in.Jti {
		return &user.AccessTokenJTICheckResp{Valid: false}, nil
	}
	return &user.AccessTokenJTICheckResp{Valid: true}, nil
}
