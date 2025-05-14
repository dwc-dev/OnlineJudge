package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"encoding/json"

	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCompetitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompetitionLogic {
	return &UpdateCompetitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCompetitionLogic) UpdateCompetition(req *types.UpdateCompetitionReq) (resp *types.UpdateCompetitionResp, err error) {
	questionsBytes, err := json.Marshal(req.Questions)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CompetitionRpc.UpdateCompetition(l.ctx, &competition.UpdateCompetitionReq{
		Competition: &competition.CompetitionInfo{
			Id:               req.Id,
			Name:             req.Name,
			Description:      req.Description,
			StartTime:        req.StartTime,
			EndTime:          req.EndTime,
			Questions:        string(questionsBytes),
			Type:             req.CompetitionType,
			PasswordRequired: req.PasswordRequired,
			Password:         req.Password,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateCompetitionResp{}, nil
}
