package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompetitionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionInfoLogic {
	return &GetCompetitionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompetitionInfoLogic) GetCompetitionInfo(req *types.GetCompetitionInfoReq) (resp *types.GetCompetitionInfoResp, err error) {
	rpcResp, err := l.svcCtx.CompetitionRpc.GetCompetitionInfo(l.ctx, &competition.GetCompetitionInfoReq{
		Id:    req.CompetitionId,
		Admin: false,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetCompetitionInfoResp{
		CompetitionInfo: types.CompetitionInfo{
			Id:               rpcResp.Competition.Id,
			Name:             rpcResp.Competition.Name,
			Description:      rpcResp.Competition.Description,
			StartTime:        rpcResp.Competition.StartTime,
			EndTime:          rpcResp.Competition.EndTime,
			CompetitionType:  rpcResp.Competition.Type,
			PasswordRequired: rpcResp.Competition.PasswordRequired,
		},
	}
	return resp, nil
}
