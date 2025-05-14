package competition

import (
	"backend/gateway/internal/logic/competition"
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"net/http"

	"backend/common/errors"
	"backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateCompetitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCompetitionReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r.Context(), w, nil, errors.InvalidParams)
			return
		}

		l := competition.NewUpdateCompetitionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCompetition(&req)
		response.Response(r.Context(), w, resp, err)
	}
}
