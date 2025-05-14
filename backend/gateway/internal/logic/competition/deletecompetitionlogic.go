package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCompetitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompetitionLogic {
	return &DeleteCompetitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCompetitionLogic) DeleteCompetition(req *types.DeleteCompetitionReq) (resp *types.DeleteCompetitionResp, err error) {
	_, err = l.svcCtx.CompetitionRpc.DeleteCompetition(l.ctx, &competition.DeleteCompetitionReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteCompetitionResp{}, nil
}
