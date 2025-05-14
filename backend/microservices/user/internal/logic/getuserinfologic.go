package logic

import (
	"context"

	"backend/microservices/user/internal/svc"
	"backend/microservices/user/pb/user"

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
	u, err := l.svcCtx.UserDao.GetUserInfoById(l.ctx, in.UserId, in.Col)
	if err != nil {
		return nil, err
	}
	userProfile := ""
	if u.UserProfile != nil {
		userProfile = *u.UserProfile
	}
	return &user.GetUserInfoResp{
		UserInfo: &user.UserInfo{
			UserId:        in.UserId,
			UserName:      u.UserName,
			UserEmail:     u.UserEmail,
			UserAvatarUrl: u.UserAvatarURL,
			UserProfile:   userProfile,
			UserRole:      u.UserRole,
			CreateAt:      u.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:      u.UpdateAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
