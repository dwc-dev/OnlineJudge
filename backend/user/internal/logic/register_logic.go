package logic

import (
	"context"

	"user/internal/db/model/account"
	"user/internal/response"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	err = l.svcCtx.UsersModel.CheckUserName(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.UsersModel.CheckUserEmail(l.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, response.ServerError
	}
	err = l.svcCtx.UsersModel.CreateNewUser(l.ctx, &account.Users{
		UserName:     req.Username,
		UserEmail:    req.Email,
		UserPassword: string(hashedPassword),
		UserRole:     "user",
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
