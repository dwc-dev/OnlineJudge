package middleware

import (
	"backend/common/errors"
	"backend/common/response"
	"net/http"
)

type AdminAuthMiddleware struct {
}

func NewAdminAuthMiddleware() *AdminAuthMiddleware {
	return &AdminAuthMiddleware{}
}

func (m *AdminAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole, ok := r.Context().Value("user_role").(string)
		if !ok {
			response.Response(r.Context(), w, nil, errors.AdminAuthFailed)
			return
		}
		if userRole != "admin" {
			response.Response(r.Context(), w, nil, errors.AdminAuthFailed)
			return
		}
		next(w, r)
	}
}
