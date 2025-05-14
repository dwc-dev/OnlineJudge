package jwt

import (
	"backend/common/errors/rpcerrors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID   uint64 `json:"user_id"`
	UserRole string `json:"user_role"`
	JTI      string `json:"jti"`
	Type     string `json:"type"`
	jwt.RegisteredClaims
}

// 生成AccessToken
func GenerateAccessToken(secretKey string, userID uint64, userRole string, jti string, issuedAt time.Time, expiresAt time.Time) (string, error) {
	// 创建Claims
	claims := TokenClaims{
		UserID:   userID,
		UserRole: userRole,
		JTI:      jti,
		Type:     "access_token",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(issuedAt),  // 设置签发时间
			ExpiresAt: jwt.NewNumericDate(expiresAt), // 设置过期时间
		},
	}

	// 创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名Token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", rpcerrors.GenerateJWTError
	}

	return tokenString, nil
}

// 生成RefreshToken
func GenerateRefreshToken(secretKey string, userID uint64, userRole string, jti string, issuedAt time.Time, expiresAt time.Time) (string, error) {
	claims := TokenClaims{
		UserID:   userID,
		UserRole: userRole,
		JTI:      jti,
		Type:     "refresh_token",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(issuedAt),  // 设置签发时间
			ExpiresAt: jwt.NewNumericDate(expiresAt), // 设置过期时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", rpcerrors.GenerateJWTError
	}

	return tokenString, nil
}
