package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"encoding/json"

	"backend/microservices/competition/competitionclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetCompetitionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetCompetitionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetCompetitionListLogic {
	return &AdminGetCompetitionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetCompetitionListLogic) AdminGetCompetitionList(req *types.AdminGetCompetitionListReq) (resp *types.AdminGetCompetitionListResp, err error) {
	rpcResp, err := l.svcCtx.CompetitionRpc.GetCompetitionList(l.ctx, &competitionclient.GetCompetitionListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Filter:   req.Filter,
	})
	if err != nil {
		return nil, err
	}
	compList := make([]*types.CompetitionInfo, len(rpcResp.CompetitionList))
	for i, comp := range rpcResp.CompetitionList {
		questions := make([]*types.CompetitionQuestion, 0)
		err = json.Unmarshal([]byte(comp.Questions), &questions)
		if err != nil {
			return nil, err
		}
		compList[i] = &types.CompetitionInfo{
			Id:               comp.Id,
			Name:             comp.Name,
			Description:      comp.Description,
			StartTime:        comp.StartTime,
			EndTime:          comp.EndTime,
			Questions:        questions,
			CompetitionType:  comp.Type,
			PasswordRequired: comp.PasswordRequired,
		}
	}
	return &types.AdminGetCompetitionListResp{
		CompetitionList: compList,
		Total:           rpcResp.Total,
		Page:            rpcResp.Page,
		PageSize:        rpcResp.PageSize,
	}, nil
}
