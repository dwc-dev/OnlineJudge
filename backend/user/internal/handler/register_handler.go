package handler

import (
	"net/http"

	"user/internal/logic"
	"user/internal/response"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r.Context(), w, nil, response.InvalidParams) // 封装统一响应
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Response(r.Context(), w, resp, err) // 封装统一响应
	}
}
