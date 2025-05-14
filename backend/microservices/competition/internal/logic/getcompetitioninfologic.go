package logic

import (
	"context"

	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCompetitionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionInfoLogic {
	return &GetCompetitionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompetitionInfoLogic) GetCompetitionInfo(in *competition.GetCompetitionInfoReq) (*competition.GetCompetitionInfoResp, error) {
	comp, err := l.svcCtx.CompetitionDao.GetCompetitionById(in.Id)
	if err != nil {
		return nil, err
	}
	if in.Admin {
		return &competition.GetCompetitionInfoResp{Competition: &competition.CompetitionInfo{
			Id:               comp.ID,
			Name:             comp.Name,
			Description:      comp.Description,
			StartTime:        comp.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:          comp.EndTime.Format("2006-01-02 15:04:05"),
			Questions:        comp.Questions,
			Type:             comp.Type,
			PasswordRequired: comp.Password != nil,
		}}, nil
	}
	return &competition.GetCompetitionInfoResp{Competition: &competition.CompetitionInfo{
		Id:               comp.ID,
		Name:             comp.Name,
		Description:      comp.Description,
		StartTime:        comp.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:          comp.EndTime.Format("2006-01-02 15:04:05"),
		Questions:        "",
		Type:             comp.Type,
		PasswordRequired: comp.Password != nil,
	}}, nil
}
