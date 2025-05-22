package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"backend/code-sandbox/types"
	"backend/microservices/judge/internal/db/model"
	"backend/microservices/judge/internal/svc"
	"backend/microservices/judge/pb/judge"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type JudgeCase struct {
	Input  string `json:"input"`  // 输入
	Output string `json:"output"` // 输出
}

type JudgeConfig struct {
	MemoryLimitMiB     int64 `json:"memory_limit_mib"`    // 内存限制，单位为MB
	StackLimitMiB      int64 `json:"stack_limit_mib"`     // 栈限制，单位为MB
	TimeoutMillisecond int   `json:"timeout_millisecond"` // 超时限制，单位为毫秒（Milliseconds）
}

type ExecDetailToDB struct {
	Output           string `json:"output"`
	TimeMilliseconds int64  `json:"time_milliseconds"`
	MemoryUsage      int64  `json:"memory_usage_kb"`
	Timeout          bool   `json:"timeout"`
	MemoryOut        bool   `json:"memory_out"`
	RuntimeError     bool   `json:"runtime_error"`
	StackOverflow    bool   `json:"stack_overflow"`
	Accepted         bool   `json:"accepted"`
}

type ResultToDB struct {
	ExecDetails        []*ExecDetailToDB `json:"exec_details"`
	CompileError       bool              `json:"compile_error"`
	CompileErrorOutput string            `json:"compile_error_output"`
	Accepted           bool              `json:"accepted"`
}

func NewJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeLogic {
	return &JudgeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeLogic) Judge(in *judge.JudgeReq) (*judge.JudgeResp, error) {
	judgeCaseRes, err := l.svcCtx.QuestionRpc.GetQuestionJudgeCase(l.ctx, &question.GetQuestionJudgeCaseReq{
		Id: in.QuestionId,
	})
	if err != nil {
		return nil, err
	}
	judgeCaseJSON := judgeCaseRes.JudgeCase
	judgeConfigRes, err := l.svcCtx.QuestionRpc.GetQuestionJudgeConfig(l.ctx, &question.GetQuestionJudgeConfigReq{
		Id: in.QuestionId,
	})
	if err != nil {
		return nil, err
	}
	judgeConfigJSON := judgeConfigRes.JudgeConfig

	fmt.Println("judgeCaseJSON:", judgeCaseJSON)
	fmt.Println("judgeConfigJSON:", judgeConfigJSON)

	var judgeCases []JudgeCase
	var judgeConfig JudgeConfig
	err = json.Unmarshal([]byte(judgeCaseJSON), &judgeCases)
	if err != nil {
		fmt.Println("error1:", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(judgeConfigJSON), &judgeConfig)
	if err != nil {
		fmt.Println("error2:", err)
		return nil, err
	}

	var inputs, outputs []string
	for _, jc := range judgeCases {
		inputs = append(inputs, jc.Input)
		outputs = append(outputs, jc.Output)
	}

	result, err := l.svcCtx.Sandbox.RunCode(l.ctx, &types.RunConfig{
		Language:           in.Language,
		Code:               in.Code,
		InputList:          inputs,
		MemoryLimitMiB:     judgeConfig.MemoryLimitMiB,
		StackLimitMiB:      judgeConfig.StackLimitMiB,
		TimeoutMillisecond: judgeConfig.TimeoutMillisecond,
	})
	if err != nil {
		return nil, err
	}

	details := make([]*judge.ExecDetail, len(result.ExecDetails))
	detailsToDB := make([]*ExecDetailToDB, len(result.ExecDetails))
	allAccepted := false
	acceptedCount := 0
	for i, execDetail := range result.ExecDetails {
		// 去掉输出结果末尾的回车符
		execDetail.Output = strings.TrimRight(execDetail.Output, "\n")
		details[i] = &judge.ExecDetail{
			Output:           execDetail.Output,
			TimeMilliseconds: execDetail.TimeMilliseconds,
			MemoryUsage:      execDetail.MemoryUsage,
			Timeout:          execDetail.Timeout,
			MemoryOut:        execDetail.MemoryOut,
			RuntimeError:     execDetail.RuntimeError,
			StackOverflow:    execDetail.StackOverflow,
			Accepted:         execDetail.Output == outputs[i],
		}
		detailsToDB[i] = &ExecDetailToDB{
			Output:           execDetail.Output,
			TimeMilliseconds: execDetail.TimeMilliseconds,
			MemoryUsage:      execDetail.MemoryUsage,
			Timeout:          execDetail.Timeout,
			MemoryOut:        execDetail.MemoryOut,
			RuntimeError:     execDetail.RuntimeError,
			StackOverflow:    execDetail.StackOverflow,
			Accepted:         execDetail.Output == outputs[i],
		}
		if details[i].Accepted {
			acceptedCount++
		}
	}
	if acceptedCount == len(details) && !result.CompileError {
		allAccepted = true
	}

	var resultToDB ResultToDB
	resultToDB.ExecDetails = detailsToDB
	resultToDB.CompileError = result.CompileError
	resultToDB.CompileErrorOutput = result.CompileErrorOutput
	resultToDB.Accepted = allAccepted
	resultJSON, err := json.Marshal(resultToDB)
	if err != nil {
		return nil, err
	}

	// 记录到数据库
	judgeRecord := &model.Judge{
		QuestionID: in.QuestionId,
		UserID:     in.UserId,
		Language:   in.Language,
		Code:       in.Code,
		ExecResult: string(resultJSON),
		Accepted:   allAccepted,
		JudgeType:  in.JudgeType,
	}
	err = l.svcCtx.JudgeDao.CreateJudge(l.ctx, judgeRecord)
	if err != nil {
		return nil, err
	}

	// 添加提交数量
	_, err = l.svcCtx.QuestionRpc.AddSubmitNum(l.ctx, &question.AddSubmitNumReq{
		Id: in.QuestionId,
	})
	if err != nil {
		return nil, err
	}

	if allAccepted {
		// 添加通过数量
		_, err = l.svcCtx.QuestionRpc.AddAcceptedNum(l.ctx, &question.AddAcceptedNumReq{
			Id: in.QuestionId,
		})
		if err != nil {
			return nil, err
		}
	}

	res := &judge.JudgeResp{
		ExecDetails:        details,
		CompileError:       result.CompileError,
		CompileErrorOutput: result.CompileErrorOutput,
		JudgeId:            judgeRecord.ID,
	}
	return res, nil
}
