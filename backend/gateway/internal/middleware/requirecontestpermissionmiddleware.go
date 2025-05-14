package middleware

import (
	"backend/common/errors"
	"backend/common/response"
	"net/http"
)

type RequireContestPermissionMiddleware struct {
}

func NewRequireContestPermissionMiddleware() *RequireContestPermissionMiddleware {
	return &RequireContestPermissionMiddleware{}
}

func (m *RequireContestPermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hasPermission := r.Context().Value("has_permission").(bool)
		if !hasPermission {
			response.Response(r.Context(), w, nil, errors.CompetitionNotAttend)
			return
		}
		next(w, r)
	}
}
