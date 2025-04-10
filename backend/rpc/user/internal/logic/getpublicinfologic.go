package logic

import (
	"context"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicInfoLogic {
	return &GetPublicInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPublicInfoLogic) GetPublicInfo(in *user.GetPublicInfoReq) (*user.GetPublicInfoResp, error) {
	u, err := l.svcCtx.UserDao.GetUserInfoById(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &user.GetPublicInfoResp{
		UserId:        in.UserId,
		UserName:      u.UserName,
		UserAvatarUrl: u.UserAvatarURL,
		UserProfile:   u.UserProfile,
	}, nil
}
