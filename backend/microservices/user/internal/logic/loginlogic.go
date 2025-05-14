package logic

import (
	"context"
	"time"

	"backend/microservices/user/internal/svc"
	"backend/microservices/user/internal/utils/jwt"
	"backend/microservices/user/pb/user"

	"github.com/google/uuid"
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
	issuedAt := time.Now()                               // 签发时间
	accessExpiresAt := issuedAt.Add(time.Minute * 15)    // 过期时间
	refreshExpiresAt := issuedAt.Add(time.Hour * 24 * 7) // 过期时间
	accessJTI := uuid.New().String()                     // AccessToken唯一标识
	refreshJTI := uuid.New().String()                    // RefreshToken唯一标识
	accessToken, err := jwt.GenerateAccessToken(l.svcCtx.Config.JWT.SecretKey, userId, userRole, accessJTI, issuedAt, accessExpiresAt)
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwt.GenerateRefreshToken(l.svcCtx.Config.JWT.SecretKey, userId, userRole, refreshJTI, issuedAt, refreshExpiresAt)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.SetAccessTokenJTI(l.ctx, userId, accessJTI)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.SetRefreshTokenJTI(l.ctx, userId, refreshJTI)
	if err != nil {
		return nil, err
	}
	return &user.LoginResp{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
