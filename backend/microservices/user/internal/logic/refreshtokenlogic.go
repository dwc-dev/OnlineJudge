package logic

import (
	"context"
	"errors"
	"time"

	"backend/microservices/user/internal/svc"
	"backend/microservices/user/pb/user"

	myjwt "backend/microservices/user/internal/utils/jwt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *user.RefreshTokenReq) (*user.RefreshTokenResp, error) {
	// 解析token
	token, err := jwt.Parse(in.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否为HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(l.svcCtx.Config.JWT.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("refresh token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("refresh token is invalid")
	}

	userID := uint64(claims["user_id"].(float64))
	refreshJTI := claims["jti"].(string)
	tokenType := claims["type"].(string)
	userRole := claims["user_role"].(string)

	storedRefreshJTI, err := l.svcCtx.RedisClient.GetRefreshTokenJTI(l.ctx, userID)
	if err != nil {
		return nil, err
	}
	if storedRefreshJTI != refreshJTI || tokenType != "refresh_token" {
		return nil, errors.New("refresh token is invalid")
	}

	issuedAt := time.Now()                               // 签发时间
	accessExpiresAt := issuedAt.Add(time.Minute * 15)    // AccessToken过期时间
	refreshExpiresAt := issuedAt.Add(time.Hour * 24 * 7) // RefreshToken过期时间
	newAccessJTI := uuid.New().String()                  // AccessToken唯一标识
	newRefreshJTI := uuid.New().String()                 // RefreshToken唯一标识

	accessToken, err := myjwt.GenerateAccessToken(l.svcCtx.Config.JWT.SecretKey, userID, userRole, newAccessJTI, issuedAt, accessExpiresAt)
	if err != nil {
		return nil, err
	}
	refreshToken, err := myjwt.GenerateRefreshToken(l.svcCtx.Config.JWT.SecretKey, userID, userRole, newRefreshJTI, issuedAt, refreshExpiresAt)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.RedisClient.SetAccessTokenJTI(l.ctx, userID, newAccessJTI)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.SetRefreshTokenJTI(l.ctx, userID, newRefreshJTI)
	if err != nil {
		return nil, err
	}

	return &user.RefreshTokenResp{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
