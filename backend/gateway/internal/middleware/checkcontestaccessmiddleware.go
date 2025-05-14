package middleware

import (
	"backend/common/errors"
	"backend/common/response"
	"backend/microservices/competition/competitionclient"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type CheckContestAccessMiddleware struct {
	competitionRpc competitionclient.Competition
}

func NewCheckContestAccessMiddleware(competitionRpc competitionclient.Competition) *CheckContestAccessMiddleware {
	return &CheckContestAccessMiddleware{competitionRpc: competitionRpc}
}

func (m *CheckContestAccessMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 复制 Body 内容，避免读取后无法被下游读取
		var bodyCopy bytes.Buffer
		tee := io.TeeReader(r.Body, &bodyCopy)
		var reqBody struct {
			CompetitionId uint64 `json:"competition_id"`
		}
		if err := json.NewDecoder(tee).Decode(&reqBody); err != nil {
			response.Response(r.Context(), w, nil, errors.InvalidParams)
			return
		}
		// 恢复 r.Body 给下游使用
		r.Body = io.NopCloser(&bodyCopy)
		resp, err := m.competitionRpc.ContestAccessCheck(r.Context(), &competitionclient.ContestAccessCheckReq{
			CompetitionId: reqBody.CompetitionId,
			UserId:        r.Context().Value("user_id").(uint64),
		})
		if err != nil {
			response.Response(r.Context(), w, nil, errors.ServerError)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "has_permission", resp.HasPermission)
		ctx = context.WithValue(ctx, "has_started", resp.HasStarted)
		ctx = context.WithValue(ctx, "is_running", resp.IsRunning)
		next(w, r.WithContext(ctx))
	}
}
