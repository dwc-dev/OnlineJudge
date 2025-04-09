package judge

import (
	"context"
	"fmt"

	"backend/api/gateway/internal/svc"
	"backend/api/gateway/internal/types"
	"backend/rpc/judge/pb/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeLogic {
	return &JudgeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JudgeLogic) Judge(req *types.JudgeReq) (resp *types.JudgeResp, err error) {
	judgeResp, err := l.svcCtx.JudgeRpc.Judge(l.ctx, &judge.JudgeReq{
		Code:       req.Code,
		Language:   req.Language,
		QuestionId: req.QuestionId,
	})
	if err != nil {
		fmt.Println("------------------------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------------------------")
		return nil, err
	}
	// 将 []*judge.ExecResult 转换为 []*types.ExecResult
	results := make([]*types.ExecResult, len(judgeResp.Results))
	for i, result := range judgeResp.Results {
		results[i] = &types.ExecResult{
			Output:           result.Output,
			TimeMilliseconds: result.TimeMilliseconds,
			MemoryUsage:      result.MemoryUsage,
			Accepted:         result.Accepted,
		}
	}
	return &types.JudgeResp{
		Results: results,
	}, nil
}
