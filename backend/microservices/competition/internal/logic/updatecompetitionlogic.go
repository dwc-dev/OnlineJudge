package logic

import (
	"context"
	"time"

	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/internal/utils/db/models"
	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type UpdateCompetitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompetitionLogic {
	return &UpdateCompetitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCompetitionLogic) UpdateCompetition(in *competition.UpdateCompetitionReq) (*competition.UpdateCompetitionResp, error) {
	comp, err := l.svcCtx.CompetitionDao.GetCompetitionById(in.Competition.Id)
	if err != nil {
		return nil, err
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, err
	}
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", in.Competition.StartTime, location)
	if err != nil {
		return nil, err
	}
	startTime = startTime.UTC()
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", in.Competition.EndTime, location)
	if err != nil {
		return nil, err
	}
	endTime = endTime.UTC()
	var passwordVersion uint = comp.PasswordVersion
	var password *string = comp.Password
	if in.Competition.Password != "" && in.Competition.PasswordRequired {
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(in.Competition.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		passwordString := string(hashedPasswordBytes)
		password = &passwordString
		passwordVersion++
	}
	if !in.Competition.PasswordRequired {
		password = nil
	}
	err = l.svcCtx.CompetitionDao.UpdateCompetition(&models.Competition{
		ID:              in.Competition.Id,
		Name:            in.Competition.Name,
		Description:     in.Competition.Description,
		StartTime:       startTime,
		EndTime:         endTime,
		Questions:       in.Competition.Questions,
		Type:            in.Competition.Type,
		Password:        password,
		PasswordVersion: passwordVersion,
	})
	if err != nil {
		return nil, err
	}
	return &competition.UpdateCompetitionResp{}, nil
}
