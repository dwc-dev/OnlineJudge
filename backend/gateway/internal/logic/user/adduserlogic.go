package user

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	_, err = l.svcCtx.UserRpc.AddUser(l.ctx, &user.AddUserReq{
		UserInfo: &user.UserInfo{
			UserId:        req.UserInfo.UserId,
			UserName:      req.UserInfo.UserName,
			UserEmail:     req.UserInfo.UserEmail,
			UserAvatarUrl: req.UserInfo.UserAvatarUrl,
			UserProfile:   req.UserInfo.UserProfile,
			UserRole:      req.UserInfo.UserRole,
		},
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddUserResp{}, nil
}
