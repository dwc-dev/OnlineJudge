package user

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicInfoLogic {
	return &GetPublicInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicInfoLogic) GetPublicInfo(req *types.GetPublicInfoReq) (resp *types.GetPublicInfoResp, err error) {
	rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		UserId: req.UserId,
		Col: []string{
			"id",
			"name",
			"avatar_url",
			"profile",
			"role",
			"create_at",
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.GetPublicInfoResp{
		UserInfo: types.UserInfo{
			UserId:        rpcResp.UserInfo.UserId,
			UserName:      rpcResp.UserInfo.UserName,
			UserAvatarUrl: rpcResp.UserInfo.UserAvatarUrl,
			UserProfile:   rpcResp.UserInfo.UserProfile,
			UserRole:      rpcResp.UserInfo.UserRole,
			CreateAt:      rpcResp.UserInfo.CreateAt,
		},
	}, nil
}
