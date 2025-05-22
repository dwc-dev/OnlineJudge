package logic

import (
	"context"
	"regexp"

	"backend/common/errors/rpcerrors"
	"backend/microservices/user/internal/svc"
	"backend/microservices/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserPasswordLogic) UpdateUserPassword(in *user.UpdateUserPasswordReq) (*user.UpdateUserPasswordResp, error) {

	// 检查旧密码是否正确
	oldPassword, err := l.svcCtx.UserDao.GetUserPasswordById(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(in.OldPassword))
	if err != nil {
		return nil, rpcerrors.OldPasswordError
	}

	// 检查新密码是否与旧密码相同
	if in.OldPassword == in.NewPassword {
		return nil, rpcerrors.NewPasswordError
	}

	// 检查新密码长度是否符合要求
	if len(in.NewPassword) < 8 || len(in.NewPassword) > 30 {
		return nil, rpcerrors.NewPasswordError
	}

	// 检查新密码是否只包含允许的字符
	if !regexp.MustCompile(`^[a-zA-Z0-9~!@#$%^&*()_+]+$`).MatchString(in.NewPassword) {
		return nil, rpcerrors.NewPasswordError
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, rpcerrors.ServerError
	}

	// 更新密码
	err = l.svcCtx.UserDao.UpdateUserPassword(l.ctx, in.UserId, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return &user.UpdateUserPasswordResp{}, nil
}
