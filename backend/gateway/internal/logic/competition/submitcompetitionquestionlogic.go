package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/competition/competitionclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitCompetitionQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitCompetitionQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitCompetitionQuestionLogic {
	return &SubmitCompetitionQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitCompetitionQuestionLogic) SubmitCompetitionQuestion(req *types.SubmitCompetitionQuestionReq) (resp *types.SubmitCompetitionQuestionResp, err error) {
	userId := l.ctx.Value("user_id").(uint64)
	rpcResp, err := l.svcCtx.CompetitionRpc.SubmitCompetitionQuestion(l.ctx, &competitionclient.SubmitCompetitionQuestionReq{
		UserId:        userId,
		CompetitionId: req.CompetitionId,
		Qid:           req.Qid,
		Code:          req.Code,
		Language:      req.Language,
	})
	if err != nil {
		return nil, err
	}
	execDetails := make([]*types.ExecDetail, len(rpcResp.JudgeResult.ExecDetails))
	for i, execDetail := range rpcResp.JudgeResult.ExecDetails {
		execDetails[i] = &types.ExecDetail{
			Output:           execDetail.Output,
			TimeMilliseconds: execDetail.TimeMilliseconds,
			MemoryUsage:      execDetail.MemoryUsage,
			Timeout:          execDetail.Timeout,
			MemoryOut:        execDetail.MemoryOut,
			RuntimeError:     execDetail.RuntimeError,
			StackOverflow:    execDetail.StackOverflow,
			Accepted:         execDetail.Accepted,
		}
	}
	return &types.SubmitCompetitionQuestionResp{
		JudgeResult: types.JudgeResult{
			ExecDetails:        execDetails,
			CompileError:       rpcResp.JudgeResult.CompileError,
			CompileErrorOutput: rpcResp.JudgeResult.CompileErrorOutput,
		},
	}, nil
}
