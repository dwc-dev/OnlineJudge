package user

import (
	"context"

	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePersonalInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonalInfoLogic {
	return &UpdatePersonalInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePersonalInfoLogic) UpdatePersonalInfo(req *types.UpdatePersonalInfoReq) (resp *types.UpdatePersonalInfoResp, err error) {
	userID := l.ctx.Value("user_id").(uint64)
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &user.UpdateUserReq{
		UserInfo: &user.UserInfo{
			UserId:      userID,
			UserName:    req.UserInfo.UserName,
			UserProfile: req.UserInfo.UserProfile,
		},
		AvatarBase64: req.AvatarBase64,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdatePersonalInfoResp{}, nil
}
