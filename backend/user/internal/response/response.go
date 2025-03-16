package response

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Result struct {
	Code int    `json:"code"`
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

func fail(e *RespError) *Result {
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
		case *RespError:
			httpx.WriteJsonCtx(ctx, w, e.HttpCode, fail(e))
		default:
			httpx.WriteJsonCtx(ctx, w, http.StatusInternalServerError, fail(ServerError))
		}
	} else {
		httpx.OkJsonCtx(ctx, w, success(v))
	}
}
