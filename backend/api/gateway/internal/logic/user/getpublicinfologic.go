package user

import (
	"context"

	"backend/api/gateway/internal/svc"
	"backend/api/gateway/internal/types"
	"backend/rpc/user/pb/user"

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
	r, err := l.svcCtx.UserRpc.GetPublicInfo(l.ctx, &user.GetPublicInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetPublicInfoResp{
		UserId:        r.UserId,
		UserName:      r.UserName,
		UserAvatarUrl: r.UserAvatarUrl,
		UserProfile:   r.UserProfile,
	}, nil
}
