package user

import (
	"context"

	"backend/api/gateway/internal/svc"
	"backend/api/gateway/internal/types"
	"backend/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	userInfo, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		UserId: uint64(l.ctx.Value("user_id").(float64)),
	})
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		UserId:        userInfo.UserId,
		UserName:      userInfo.UserName,
		UserEmail:     userInfo.UserEmail,
		UserAvatarUrl: userInfo.UserAvatarUrl,
		UserProfile:   userInfo.UserProfile,
		UserRole:      userInfo.UserRole,
	}, nil
}
