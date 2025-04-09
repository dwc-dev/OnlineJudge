package jwt

import (
	"backend/common/errors/rpcerrors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 定义一个结构体来存储JWT的Claims
type CustomClaims struct {
	UserID   uint64 `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}

// 生成JWT
func GenerateJWT(secretKey string, userID uint64, userRole string, issuedAt time.Time, expirationTime time.Time) (string, error) {
	// 创建Claims
	claims := CustomClaims{
		UserID:   userID,
		UserRole: userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(issuedAt),       // 设置签发时间
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 设置过期时间
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
