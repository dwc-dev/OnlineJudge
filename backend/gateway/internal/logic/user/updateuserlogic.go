package user

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &user.UpdateUserReq{
		UserInfo: &user.UserInfo{
			UserId:      req.UserInfo.UserId,
			UserName:    req.UserInfo.UserName,
			UserEmail:   req.UserInfo.UserEmail,
			UserProfile: req.UserInfo.UserProfile,
			UserRole:    req.UserInfo.UserRole,
		},
		Password:     req.Password,
		AvatarBase64: req.AvatarBase64,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateUserResp{}, nil
}
