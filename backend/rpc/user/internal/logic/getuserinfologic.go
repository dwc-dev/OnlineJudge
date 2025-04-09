package logic

import (
	"context"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	u, err := l.svcCtx.UserDao.GetUserInfoById(l.ctx, uint(in.UserId))
	if err != nil {
		return nil, err
	}
	return &user.GetUserInfoResp{
		UserId:        uint64(u.ID),
		UserName:      u.UserName,
		UserEmail:     u.UserEmail,
		UserAvatarUrl: u.UserAvatarURL,
		UserProfile:   u.UserProfile,
		UserRole:      u.UserRole,
	}, nil
}
