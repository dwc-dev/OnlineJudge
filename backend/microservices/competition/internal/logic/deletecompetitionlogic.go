package logic

import (
	"context"

	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCompetitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompetitionLogic {
	return &DeleteCompetitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCompetitionLogic) DeleteCompetition(in *competition.DeleteCompetitionReq) (*competition.DeleteCompetitionResp, error) {
	err := l.svcCtx.CompetitionDao.DeleteCompetition(in.Id)
	if err != nil {
		return nil, err
	}
	return &competition.DeleteCompetitionResp{}, nil
}
