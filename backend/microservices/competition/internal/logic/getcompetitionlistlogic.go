package logic

import (
	"context"

	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCompetitionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionListLogic {
	return &GetCompetitionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompetitionListLogic) GetCompetitionList(in *competition.GetCompetitionListReq) (*competition.GetCompetitionListResp, error) {
	compList, total, err := l.svcCtx.CompetitionDao.GetCompetitionList(in.Page, in.PageSize, in.Filter, in.Col)
	if err != nil {
		return nil, err
	}
	compListResp := make([]*competition.CompetitionInfo, len(compList))
	for i, comp := range compList {
		compListResp[i] = &competition.CompetitionInfo{
			Id:               comp.ID,
			Name:             comp.Name,
			Description:      comp.Description,
			StartTime:        comp.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:          comp.EndTime.Format("2006-01-02 15:04:05"),
			Questions:        comp.Questions,
			Type:             comp.Type,
			PasswordRequired: comp.Password != nil,
		}
	}
	return &competition.GetCompetitionListResp{
		CompetitionList: compListResp,
		Total:           total,
		Page:            in.Page,
		PageSize:        in.PageSize,
	}, nil
}
