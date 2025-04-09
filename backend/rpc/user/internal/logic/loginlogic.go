package logic

import (
	"context"
	"time"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/internal/utils/jwt"
	"backend/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	userId, userRole, err := l.svcCtx.UserDao.CompareUserPassword(l.ctx, in.Email, in.Password)
	if err != nil {
		return nil, err

	}
	issuedAt := time.Now()                                                                  // 签发时间
	expirationTime := issuedAt.Add(time.Duration(l.svcCtx.Config.JWT.Expire) * time.Second) // 过期时间
	token, err := jwt.GenerateJWT(l.svcCtx.Config.JWT.Secret, userId, userRole, issuedAt, expirationTime)
	if err != nil {
		return nil, err
	}
	return &user.LoginResp{Token: token}, nil
}
