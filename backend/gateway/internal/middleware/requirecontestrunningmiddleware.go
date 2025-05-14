package middleware

import (
	"backend/common/errors"
	"backend/common/response"
	"net/http"
)

type RequireContestRunningMiddleware struct {
}

func NewRequireContestRunningMiddleware() *RequireContestRunningMiddleware {
	return &RequireContestRunningMiddleware{}
}

func (m *RequireContestRunningMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isRunning := r.Context().Value("is_running").(bool)
		if !isRunning {
			response.Response(r.Context(), w, nil, errors.CompetitionNotRunning)
			return
		}
		next(w, r)
	}
}
