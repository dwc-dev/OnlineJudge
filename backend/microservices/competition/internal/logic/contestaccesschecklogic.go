package logic

import (
	"context"
	"time"

	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContestAccessCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContestAccessCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContestAccessCheckLogic {
	return &ContestAccessCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContestAccessCheckLogic) ContestAccessCheck(in *competition.ContestAccessCheckReq) (*competition.ContestAccessCheckResp, error) {
	comp, err := l.svcCtx.CompetitionDao.GetCompetitionById(in.CompetitionId)
	if err != nil {
		return nil, err
	}
	attendance, err := l.svcCtx.CompetitionAttendanceDao.GetCompetitionAttendance(in.CompetitionId, in.UserId)
	if err != nil {
		return nil, err
	}
	if comp.Password != nil {
		if attendance.PasswordVersion == nil || *attendance.PasswordVersion != comp.PasswordVersion {
			return &competition.ContestAccessCheckResp{
				HasPermission: false,
				HasStarted:    false,
				IsRunning:     false,
			}, nil
		}
	}
	hasStarted := false
	isRunning := false
	now := time.Now()
	if now.After(comp.StartTime) {
		hasStarted = true
	}
	if now.After(comp.StartTime) && now.Before(comp.EndTime) {
		isRunning = true
	}
	return &competition.ContestAccessCheckResp{
		HasPermission: true,
		HasStarted:    hasStarted,
		IsRunning:     isRunning,
	}, nil
}
