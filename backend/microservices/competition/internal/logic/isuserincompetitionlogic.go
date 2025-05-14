package logic

import (
	"context"
	"time"

	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUserInCompetitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsUserInCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserInCompetitionLogic {
	return &IsUserInCompetitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsUserInCompetitionLogic) IsUserInCompetition(in *competition.IsUserInCompetitionReq) (*competition.IsUserInCompetitionResp, error) {
	attendances, err := l.svcCtx.CompetitionAttendanceDao.GetCompetitionAttendancesByUserId(in.UserId)
	if err != nil {
		return nil, err
	}
	for _, attendance := range attendances {
		comp, err := l.svcCtx.CompetitionDao.GetCompetitionById(attendance.CompetitionID)
		if err != nil {
			return nil, err
		}
		if comp.Password != nil && (attendance.PasswordVersion == nil || *attendance.PasswordVersion != comp.PasswordVersion) {
			continue
		}
		now := time.Now()
		if now.After(*comp.StartTime) && now.Before(*comp.EndTime) {
			return &competition.IsUserInCompetitionResp{IsInCompetition: true}, nil
		}
	}
	return &competition.IsUserInCompetitionResp{IsInCompetition: false}, nil
}
