package middleware

import (
	"backend/common/errors"
	"backend/common/response"
	"context"
	ers "errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type JWTAuthMiddleware struct {
	accessSecret string
}

func NewJWTAuthMiddleware(accessSecret string) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{accessSecret: accessSecret}
}

func (m *JWTAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取token
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			response.Response(r.Context(), w, nil, errors.JWTInvalid)
			return
		}
		// 检查token格式是否为"Bearer <token>"
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			response.Response(r.Context(), w, nil, errors.JWTInvalid)
			return
		}
		tokenString = tokenParts[1]

		// 解析token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法是否为HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(m.accessSecret), nil
		})

		if err != nil {
			if ers.Is(err, jwt.ErrTokenExpired) {
				response.Response(r.Context(), w, nil, errors.JWTExpired)
				return
			}
			response.Response(r.Context(), w, nil, errors.JWTInvalid)
			return
		}

		// 验证token是否有效
		if !token.Valid {
			response.Response(r.Context(), w, nil, errors.JWTInvalid)
			return
		}

		//获取jwt中的payload
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Response(r.Context(), w, nil, errors.JWTInvalid)
			return
		}

		ctx := r.Context()
		for k, v := range claims {
			ctx = context.WithValue(ctx, k, v)
		}
		next(w, r.WithContext(ctx))
	}
}
