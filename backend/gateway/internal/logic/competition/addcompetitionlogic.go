package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"encoding/json"

	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCompetitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCompetitionLogic {
	return &AddCompetitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCompetitionLogic) AddCompetition(req *types.AddCompetitionReq) (resp *types.AddCompetitionResp, err error) {
	questionsBytes, err := json.Marshal(req.Questions)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CompetitionRpc.AddCompetition(l.ctx, &competition.AddCompetitionReq{
		Competition: &competition.CompetitionInfo{
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
	return &types.AddCompetitionResp{}, nil
}
