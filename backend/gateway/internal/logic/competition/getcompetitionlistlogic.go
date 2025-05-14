package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompetitionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionListLogic {
	return &GetCompetitionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompetitionListLogic) GetCompetitionList(req *types.GetCompetitionListReq) (resp *types.GetCompetitionListResp, err error) {
	rpcResp, err := l.svcCtx.CompetitionRpc.GetCompetitionList(l.ctx, &competition.GetCompetitionListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Filter:   req.Filter,
		Col:      []string{"id", "name", "start_time", "end_time", "type", "password"},
	})
	if err != nil {
		return nil, err
	}
	compList := make([]*types.CompetitionInfo, len(rpcResp.CompetitionList))
	for i, comp := range rpcResp.CompetitionList {
		compList[i] = &types.CompetitionInfo{
			Id:               comp.Id,
			Name:             comp.Name,
			StartTime:        comp.StartTime,
			EndTime:          comp.EndTime,
			CompetitionType:  comp.Type,
			PasswordRequired: comp.PasswordRequired,
		}
	}

	return &types.GetCompetitionListResp{
		CompetitionList: compList,
		Total:           rpcResp.Total,
		Page:            rpcResp.Page,
		PageSize:        rpcResp.PageSize,
	}, nil
}
