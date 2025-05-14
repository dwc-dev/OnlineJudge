package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/competition/competitionclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendCompetitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttendCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendCompetitionLogic {
	return &AttendCompetitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttendCompetitionLogic) AttendCompetition(req *types.AttendCompetitionReq) (resp *types.AttendCompetitionResp, err error) {
	userId := l.ctx.Value("user_id").(uint64) // 从jwt中读取用户id
	_, err = l.svcCtx.CompetitionRpc.AttendCompetition(l.ctx, &competitionclient.AttendCompetitionReq{
		UserId:        userId,
		CompetitionId: req.CompetitionId,
		Password:      req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.AttendCompetitionResp{}, nil
}
