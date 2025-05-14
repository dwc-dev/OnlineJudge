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

func UpdateQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateQuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r.Context(), w, nil, errors.InvalidParams)
			return
		}

		l := question.NewUpdateQuestionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateQuestion(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r.Context(), w, resp, err)
	}
}
