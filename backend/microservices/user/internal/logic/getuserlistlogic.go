package logic

import (
	"context"

	"backend/microservices/user/internal/svc"
	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *user.GetUserListReq) (*user.GetUserListResp, error) {
	users, total, err := l.svcCtx.UserDao.GetUserList(l.ctx, in.Page, in.PageSize, in.Filter)
	if err != nil {
		return nil, err
	}

	userInfos := make([]*user.UserInfo, 0)
	for _, userItem := range users {
		userInfos = append(userInfos, &user.UserInfo{
			UserId:        userItem.ID,
			UserName:      userItem.Name,
			UserEmail:     userItem.Email,
			UserAvatarUrl: userItem.AvatarURL,
			UserProfile:   userItem.Profile,
			UserRole:      userItem.Role,
			CreateAt:      userItem.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:      userItem.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &user.GetUserListResp{
		UserInfo: userInfos,
		Total:    total,
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}
