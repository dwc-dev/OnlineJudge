package logic

import (
	"context"

	"backend/common/errors/rpcerrors"
	"backend/microservices/user/internal/svc"
	"backend/microservices/user/internal/utils/db/model"
	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	userName, err := l.svcCtx.UserDao.GetUserNameById(l.ctx, in.UserInfo.UserId)
	if err != nil {
		return nil, err
	}
	userEmail, err := l.svcCtx.UserDao.GetUserEmailById(l.ctx, in.UserInfo.UserId)
	if err != nil {
		return nil, err
	}

	if userName != in.UserInfo.UserName {
		err = l.svcCtx.UserDao.CheckUserName(l.ctx, in.UserInfo.UserName)
		if err != nil {
			return nil, err
		}
	}
	if userEmail != in.UserInfo.UserEmail {
		err = l.svcCtx.UserDao.CheckUserEmail(l.ctx, in.UserInfo.UserEmail)
		if err != nil {
			return nil, err
		}
	}

	var avatar_url string

	if in.AvatarBase64 != "" {
		err = l.svcCtx.MinioClient.UploadAvatar(l.ctx, in.UserInfo.UserId, in.AvatarBase64)
		if err != nil {
			return nil, err
		}
		url := l.svcCtx.MinioClient.GetAvatarUrl(in.UserInfo.UserId)
		avatar_url = url
	}

	newInfo := &model.User{
		ID:            in.UserInfo.UserId,
		UserName:      in.UserInfo.UserName,
		UserEmail:     in.UserInfo.UserEmail,
		UserProfile:   &in.UserInfo.UserProfile,
		UserRole:      in.UserInfo.UserRole,
		UserAvatarURL: avatar_url,
	}
	if in.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, rpcerrors.ServerError
		}
		newInfo.UserPassword = string(hashedPassword)
	}
	err = l.svcCtx.UserDao.UpdateUser(l.ctx, newInfo)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserResp{}, nil
}
