package question

import (
	"backend/gateway/internal/logic/question"
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"net/http"

	"backend/common/errors"
	"backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddQuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r.Context(), w, nil, errors.InvalidParams)
			return
		}

		l := question.NewAddQuestionLogic(r.Context(), svcCtx)
		resp, err := l.AddQuestion(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r.Context(), w, resp, err)
	}
}
