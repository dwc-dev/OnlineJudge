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

func PaginationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PaginationReq
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r.Context(), w, nil, errors.InvalidParams) // 封装统一响应
			return
		}

		l := question.NewPaginationLogic(r.Context(), svcCtx)
		resp, err := l.Pagination(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r.Context(), w, resp, err) // 封装统一响应
	}
}
