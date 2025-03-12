package response

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) *Result {
	return &Result{
		Code: 0, // 成功固定为 0
		Msg:  "success",
		Data: data,
	}
}

func Fail(e *RespError) *Result {
	return &Result{
		Code: e.Code,
		Msg:  e.Msg,
		Data: nil,
	}
}
