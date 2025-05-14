package user

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/user/pb/user"

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
	rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		UserId: l.ctx.Value("user_id").(uint64),
	})
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		UserInfo: types.UserInfo{
			UserId:        rpcResp.UserInfo.UserId,
			UserName:      rpcResp.UserInfo.UserName,
			UserEmail:     rpcResp.UserInfo.UserEmail,
			UserAvatarUrl: rpcResp.UserInfo.UserAvatarUrl,
			UserProfile:   rpcResp.UserInfo.UserProfile,
			UserRole:      rpcResp.UserInfo.UserRole,
			CreateAt:      rpcResp.UserInfo.CreateAt,
			UpdateAt:      rpcResp.UserInfo.UpdateAt,
		},
	}, nil
}
