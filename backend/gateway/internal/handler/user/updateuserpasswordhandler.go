package user

import (
	"net/http"

	"backend/common/errors"
	"backend/common/response"
	"backend/gateway/internal/logic/user"
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r.Context(), w, nil, errors.InvalidParams)
			return
		}

		l := user.NewUpdateUserPasswordLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserPassword(&req)
		response.Response(r.Context(), w, resp, err)
	}
}
