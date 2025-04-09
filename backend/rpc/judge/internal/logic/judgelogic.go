package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"backend/code-sandbox/sandbox"
	"backend/rpc/judge/internal/svc"
	"backend/rpc/judge/pb/judge"
	"backend/rpc/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type JudgeCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

func NewJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeLogic {
	return &JudgeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeLogic) Judge(in *judge.JudgeReq) (*judge.JudgeResp, error) {
	fmt.Println("language:", in.Language)
	q, err := l.svcCtx.QuestionRpc.GetQuestionJudgeCase(l.ctx, &question.QuestionRequest{
		Id: in.QuestionId,
	})
	if err != nil {
		return nil, err
	}
	judgeCaseJSON := q.JudgeCase
	var judgeCases []JudgeCase
	err = json.Unmarshal([]byte(judgeCaseJSON), &judgeCases)
	if err != nil {
		return nil, err
	}
	var inputs, outputs []string
	for _, jc := range judgeCases {
		inputs = append(inputs, jc.Input)
		outputs = append(outputs, jc.Output)
	}
	fmt.Println("Inputs:", inputs)
	fmt.Println("Outputs:", outputs)
	r, err := sandbox.RunCode(in.Language, in.Code, inputs)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 将 []types.ExecResult 转换为 []*judge.ExecResult
	results := make([]*judge.ExecResult, len(r))
	for i, execResult := range r {
		// 去掉输出结果末尾的回车符
		execResult.Output = strings.TrimRight(execResult.Output, "\n")
		results[i] = &judge.ExecResult{
			Output:           execResult.Output,
			TimeMilliseconds: execResult.Time.Milliseconds(), //纳秒（Nanoseconds）转毫秒（Milliseconds）
			MemoryUsage:      int64(execResult.MemoryUsage),
			Accepted:         execResult.Output == outputs[i],
		}
	}
	res := &judge.JudgeResp{
		Results: results,
	}
	fmt.Println("------------------------------------------------------")
	fmt.Println(res)
	fmt.Println("------------------------------------------------------")
	return res, nil
}
