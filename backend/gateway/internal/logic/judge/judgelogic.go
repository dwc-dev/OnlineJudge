package judge

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/judge/pb/judge"

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
	userId := l.ctx.Value("user_id").(uint64)
	judgeResp, err := l.svcCtx.JudgeRpc.Judge(l.ctx, &judge.JudgeReq{
		Code:       req.Code,
		Language:   req.Language,
		QuestionId: req.QuestionId,
		UserId:     userId,
		JudgeType:  "normal",
	})
	if err != nil {
		return nil, err
	}

	details := make([]*types.ExecDetail, len(judgeResp.ExecDetails))
	for i, detail := range judgeResp.ExecDetails {
		details[i] = &types.ExecDetail{
			Output:           detail.Output,
			TimeMilliseconds: detail.TimeMilliseconds,
			MemoryUsage:      detail.MemoryUsage,
			Timeout:          detail.Timeout,
			MemoryOut:        detail.MemoryOut,
			RuntimeError:     detail.RuntimeError,
			StackOverflow:    detail.StackOverflow,
			Accepted:         detail.Accepted,
		}
	}
	return &types.JudgeResp{
		ExecDetails:        details,
		CompileError:       judgeResp.CompileError,
		CompileErrorOutput: judgeResp.CompileErrorOutput,
	}, nil
}
