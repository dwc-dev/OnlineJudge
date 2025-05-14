package judge

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/judge/pb/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJudgeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJudgeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJudgeListLogic {
	return &GetJudgeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJudgeListLogic) GetJudgeList(req *types.GetJudgeListReq) (resp *types.GetJudgeListResp, err error) {
	userID := l.ctx.Value("user_id").(uint64)
	rpcResp, err := l.svcCtx.JudgeRpc.GetJudgeList(l.ctx, &judge.GetJudgeListReq{
		UserId:    userID,
		Page:      req.Page,
		PageSize:  req.PageSize,
		JudgeType: "normal",
	})
	if err != nil {
		return nil, err
	}
	var judgeList []*types.JudgeInfo
	for _, judge := range rpcResp.JudgeList {
		judgeList = append(judgeList, &types.JudgeInfo{
			JudgeId:    judge.JudgeId,
			UserId:     judge.UserId,
			QuestionId: judge.QuestionId,
			Language:   judge.Language,
			Code:       judge.Code,
			ExecResult: judge.ExecResult,
			Accepted:   judge.Accepted,
			CreateAt:   judge.CreateAt,
			UpdateAt:   judge.UpdateAt,
		})
	}
	return &types.GetJudgeListResp{
		JudgeList: judgeList,
		Total:     rpcResp.Total,
		Page:      rpcResp.Page,
		PageSize:  rpcResp.PageSize,
	}, nil
}
