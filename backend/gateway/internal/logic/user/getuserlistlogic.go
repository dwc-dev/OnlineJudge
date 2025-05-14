package user

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.GetUserListResp, err error) {
	rpcResp, err := l.svcCtx.UserRpc.GetUserList(l.ctx, &user.GetUserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Filter:   req.Filter,
	})
	if err != nil {
		return nil, err
	}
	users := make([]*types.UserInfo, 0)
	for _, user := range rpcResp.UserInfo {
		users = append(users, &types.UserInfo{
			UserId:        user.UserId,
			UserName:      user.UserName,
			UserEmail:     user.UserEmail,
			UserAvatarUrl: user.UserAvatarUrl,
			UserProfile:   user.UserProfile,
			UserRole:      user.UserRole,
			CreateAt:      user.CreateAt,
			UpdateAt:      user.UpdateAt,
		})
	}
	return &types.GetUserListResp{
		UserInfo: users,
		Total:    rpcResp.Total,
		Page:     rpcResp.Page,
		PageSize: rpcResp.PageSize,
	}, nil
}
