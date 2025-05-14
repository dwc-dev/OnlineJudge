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

func SubmitCompetitionQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitCompetitionQuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r.Context(), w, nil, errors.InvalidParams)
			return
		}

		l := competition.NewSubmitCompetitionQuestionLogic(r.Context(), svcCtx)
		resp, err := l.SubmitCompetitionQuestion(&req)
		response.Response(r.Context(), w, resp, err)
	}
}
