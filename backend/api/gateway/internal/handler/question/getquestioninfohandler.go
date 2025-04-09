package question

import (
	"net/http"

	"backend/api/gateway/internal/logic/question"
	"backend/api/gateway/internal/svc"
	"backend/api/gateway/internal/types"
	"backend/common/errors"
	"backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetQuestionInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r.Context(), w, nil, errors.InvalidParams) // 封装统一响应
			return
		}

		l := question.NewGetQuestionInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetQuestionInfo(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r.Context(), w, resp, err) // 封装统一响应
	}
}
