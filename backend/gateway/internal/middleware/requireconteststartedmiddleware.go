package middleware

import (
	"backend/common/errors"
	"backend/common/response"
	"net/http"
)

type RequireContestStartedMiddleware struct {
}

func NewRequireContestStartedMiddleware() *RequireContestStartedMiddleware {
	return &RequireContestStartedMiddleware{}
}

func (m *RequireContestStartedMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hasStarted := r.Context().Value("has_started").(bool)
		if !hasStarted {
			response.Response(r.Context(), w, nil, errors.CompetitionNotStarted)
			return
		}
		next(w, r)
	}
}
