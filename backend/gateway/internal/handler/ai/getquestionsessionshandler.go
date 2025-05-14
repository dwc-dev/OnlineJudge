package ai

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	myerrors "backend/common/errors"
	"backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetQuestionSessionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetQuestionSessionsReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r.Context(), w, nil, myerrors.InvalidParams)
			return
		}
		userId, ok := r.Context().Value("user_id").(uint64)
		if !ok {
			response.Response(r.Context(), w, nil, myerrors.InvalidParams)
			return
		}
		URL := fmt.Sprintf("%s/question/sessions", svcCtx.Config.AIServiceURL)
		httpReq, err := http.NewRequest("POST", URL, nil)
		if err != nil {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		jsonBody, err := json.Marshal(struct {
			UserId     uint64 `json:"user_id"`
			QuestionId uint64 `json:"question_id"`
		}{
			UserId:     userId,
			QuestionId: req.QuestionId,
		})
		if err != nil {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		httpReq.Body = io.NopCloser(bytes.NewReader(jsonBody))
		httpReq.Header.Set("Content-Type", "application/json")
		httpClient := http.Client{
			Timeout: 0 * time.Second,
		}
		httpResp, err := httpClient.Do(httpReq)
		if err != nil {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		defer httpResp.Body.Close()
		bodyBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		var respBody struct {
			Code int             `json:"code"`
			Msg  string          `json:"msg"`
			Data []types.Session `json:"data"`
		}
		err = json.Unmarshal(bodyBytes, &respBody)
		if err != nil {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		if respBody.Code != 200 {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		response.Response(r.Context(), w, respBody.Data, nil)
	}
}
