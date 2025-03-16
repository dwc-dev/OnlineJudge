package logic

import (
	"context"
	"time"

	"user/internal/response"
	"user/internal/svc"
	"user/internal/types"
	"user/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userId, err := l.svcCtx.UsersModel.CompareUserPassword(l.ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	issuedAt := time.Now()                                                                  // 签发时间
	expirationTime := issuedAt.Add(time.Duration(l.svcCtx.Config.JWT.Expire) * time.Second) // 过期时间
	token, err := utils.GenerateJWT(l.svcCtx.Config.JWT.Secret, userId, issuedAt, expirationTime)
	if err != nil {
		return nil, response.GenerateJWTError
	}
	return &types.LoginResp{Token: token}, nil
}
