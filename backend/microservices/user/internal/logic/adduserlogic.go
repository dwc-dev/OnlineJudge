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

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserReq) (*user.AddUserResp, error) {
	// 检查用户名是否已存在
	err := l.svcCtx.UserDao.CheckUserName(l.ctx, in.UserInfo.UserName)
	if err != nil {
		return nil, err
	}
	// 检查邮箱是否已存在
	err = l.svcCtx.UserDao.CheckUserEmail(l.ctx, in.UserInfo.UserEmail)
	if err != nil {
		return nil, err
	}
	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, rpcerrors.ServerError
	}
	defaultAvatarUrl := l.svcCtx.MinioClient.GetDefaultAvatarUrl()
	defaultProfile := "这个人很懒，什么都没有留下"
	newUser := &model.User{
		UserName:      in.UserInfo.UserName,
		UserEmail:     in.UserInfo.UserEmail,
		UserRole:      in.UserInfo.UserRole,
		UserPassword:  string(hashedPassword),
		UserAvatarURL: defaultAvatarUrl,
		UserProfile:   &defaultProfile,
	}
	err = l.svcCtx.UserDao.AddUser(l.ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &user.AddUserResp{}, nil
}
