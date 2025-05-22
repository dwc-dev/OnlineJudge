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

type AddCompetitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCompetitionLogic {
	return &AddCompetitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCompetitionLogic) AddCompetition(in *competition.AddCompetitionReq) (*competition.AddCompetitionResp, error) {
	// 传入的开始时间和结束时间为 UTC+8，需要转换为 UTC 时间存入 MySQL
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
	var password *string
	var passwordVersion uint = 0
	if in.Competition.Password != "" {
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(in.Competition.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		passwordString := string(hashedPasswordBytes)
		password = &passwordString
		passwordVersion++
	}
	err = l.svcCtx.CompetitionDao.CreateCompetition(&models.Competition{
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
	return &competition.AddCompetitionResp{}, nil
}
