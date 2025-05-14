package middleware

import (
	"backend/common/errors"
	"backend/common/response"
	"backend/microservices/user/userclient"
	"context"
	ers "errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenAuthMiddleware struct {
	accessSecret string
	userRpc      userclient.User
}

func NewAccessTokenAuthMiddleware(accessSecret string, userRpc userclient.User) *AccessTokenAuthMiddleware {
	return &AccessTokenAuthMiddleware{accessSecret: accessSecret, userRpc: userRpc}
}

func (m *AccessTokenAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取token
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			response.Response(r.Context(), w, nil, errors.AccessTokenInvalid)
			return
		}
		// 检查token格式是否为"Bearer <token>"
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			response.Response(r.Context(), w, nil, errors.AccessTokenInvalid)
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
			fmt.Println("--------------------------------")
			fmt.Println("access token error")
			fmt.Println("Token:", tokenString)
			fmt.Println("URL:", r.RequestURI)
			fmt.Println("--------------------------------")
			if ers.Is(err, jwt.ErrTokenExpired) {
				response.Response(r.Context(), w, nil, errors.AccessTokenInvalid)
				return
			}
			response.Response(r.Context(), w, nil, errors.AccessTokenInvalid)
			return
		}
		// 验证token是否有效
		if !token.Valid {
			response.Response(r.Context(), w, nil, errors.AccessTokenInvalid)
			return
		}
		//获取jwt中的payload
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Response(r.Context(), w, nil, errors.AccessTokenInvalid)
			return
		}
		userID := uint64(claims["user_id"].(float64))
		userRole := claims["user_role"].(string)
		JTI := claims["jti"].(string)
		resp, err := m.userRpc.AccessTokenJTICheck(r.Context(), &userclient.AccessTokenJTICheckReq{
			UserId: userID,
			Jti:    JTI,
		})
		if err != nil {
			response.Response(r.Context(), w, nil, errors.AccessTokenInvalid)
			return
		}
		if !resp.Valid {
			response.Response(r.Context(), w, nil, errors.AccessTokenInvalid)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", userID)
		ctx = context.WithValue(ctx, "user_role", userRole)
		next(w, r.WithContext(ctx))
	}
}
