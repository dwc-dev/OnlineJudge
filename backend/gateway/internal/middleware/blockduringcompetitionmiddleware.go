package middleware

import (
	"backend/common/errors"
	"backend/common/response"
	"backend/microservices/competition/competitionclient"
	"net/http"
)

type BlockDuringCompetitionMiddleware struct {
	competitionRpc competitionclient.Competition
}

func NewBlockDuringCompetitionMiddleware(competitionRpc competitionclient.Competition) *BlockDuringCompetitionMiddleware {
	return &BlockDuringCompetitionMiddleware{competitionRpc: competitionRpc}
}

func (m *BlockDuringCompetitionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rpcResp, err := m.competitionRpc.IsUserInCompetition(r.Context(), &competitionclient.IsUserInCompetitionReq{
			UserId: r.Context().Value("user_id").(uint64),
		})
		if err != nil {
			response.Response(r.Context(), w, nil, errors.ServerError)
			return
		}
		if rpcResp.IsInCompetition {
			response.Response(r.Context(), w, nil, errors.BlockDuringCompetition)
			return
		}
		next(w, r)
	}
}
