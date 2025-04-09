package response

import (
	"backend/common/codes"
	"backend/common/errors"
	"backend/common/errors/rpcerrors"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Result struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func success(data any) *Result {
	return &Result{
		Code: 0, // 成功固定为 0
		Msg:  "success",
		Data: data,
	}
}

func fail(e *errors.Error) *Result {
	return &Result{
		Code: e.Code,
		Msg:  e.Msg,
		Data: nil,
	}
}

// 封装统一响应
func Response(ctx context.Context, w http.ResponseWriter, v any, err error) {
	if err != nil {
		switch e := err.(type) {
		case *errors.Error:
			httpx.WriteJsonCtx(ctx, w, codes.GetHTTPStatus(e.Code), fail(e))
		default:
			st, ok := rpcerrors.FromError(err)
			if ok {
				httpx.WriteJsonCtx(ctx, w,
					codes.GetHTTPStatus(st.Code),
					fail(&errors.Error{Code: st.Code, Msg: st.Message}))
			} else {
				httpx.WriteJsonCtx(ctx, w, codes.GetHTTPStatus(errors.UnknownError.Code), fail(errors.UnknownError))
			}
		}
	} else {
		httpx.OkJsonCtx(ctx, w, success(v))
	}
}
