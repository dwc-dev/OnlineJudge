package judge

import (
	"net/http"

	"backend/api/gateway/internal/logic/judge"
	"backend/api/gateway/internal/svc"
	"backend/api/gateway/internal/types"
	"backend/common/errors"
	"backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func JudgeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JudgeReq
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r.Context(), w, nil, errors.InvalidParams) // 封装统一响应
			return
		}

		l := judge.NewJudgeLogic(r.Context(), svcCtx)
		resp, err := l.Judge(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r.Context(), w, resp, err) // 封装统一响应
	}
}
