package user

import (
	"context"

	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"backend/microservices/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPasswordLogic) UpdateUserPassword(req *types.UpdateUserPasswordReq) (resp *types.UpdateUserPasswordResp, err error) {
	userID := l.ctx.Value("user_id").(uint64)
	_, err = l.svcCtx.UserRpc.UpdateUserPassword(l.ctx, &userclient.UpdateUserPasswordReq{
		UserId:      userID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateUserPasswordResp{}, nil
}
