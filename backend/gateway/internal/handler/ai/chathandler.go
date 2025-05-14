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

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r.Context(), w, nil, myerrors.InvalidParams)
			return
		}
		userId, ok := r.Context().Value("user_id").(uint64)
		if !ok {
			response.Response(r.Context(), w, nil, myerrors.InvalidParams)
			return
		}
		chatURL := fmt.Sprintf("%s/chat", svcCtx.Config.AIServiceURL)
		chatReq, err := http.NewRequest("POST", chatURL, nil)
		if err != nil {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		jsonBody, err := json.Marshal(struct {
			UserId     uint64 `json:"user_id"`
			QuestionId uint64 `json:"question_id"`
			SessionId  string `json:"session_id"`
			Message    string `json:"message"`
		}{
			UserId:     userId,
			QuestionId: req.QuestionId,
			SessionId:  req.SessionId,
			Message:    req.Message,
		})
		if err != nil {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		chatReq.Body = io.NopCloser(bytes.NewReader(jsonBody))
		chatReq.Header.Set("Content-Type", "application/json")
		client := http.Client{
			Timeout: 0 * time.Second,
		}
		chatResp, err := client.Do(chatReq)
		if err != nil {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		defer chatResp.Body.Close()
		// 设置 SSE 响应头
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			response.Response(r.Context(), w, nil, myerrors.ServerError)
			return
		}
		// 从 SSE 服务读取并转发到客户端
		buf := make([]byte, 4096)
		for {
			n, err := chatResp.Body.Read(buf)
			if err != nil {
				if err == io.EOF {
					return
				}
				logx.Error("read sse stream error:", err)
				return
			}
			_, err = w.Write(buf[:n])
			if err != nil {
				logx.Error("write to client error:", err)
				return
			}
			flusher.Flush()
		}
	}
}
