package logic

import (
	"context"

	"backend/common/errors/rpcerrors"
	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/internal/utils/db/models"
	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AttendCompetitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendCompetitionLogic {
	return &AttendCompetitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AttendCompetitionLogic) AttendCompetition(in *competition.AttendCompetitionReq) (*competition.AttendCompetitionResp, error) {
	comp, err := l.svcCtx.CompetitionDao.GetCompetitionById(in.CompetitionId)
	if err != nil {
		return nil, err
	}
	attendance, err := l.svcCtx.CompetitionAttendanceDao.GetCompetitionAttendance(in.CompetitionId, in.UserId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if attendance != nil && attendance.PasswordVersion != nil && *attendance.PasswordVersion == comp.PasswordVersion {
		return nil, rpcerrors.CompetitionAlreadyAttend
	}
	if comp.Password == nil && attendance != nil {
		return nil, rpcerrors.CompetitionAlreadyAttend
	}
	if comp.Password != nil {
		err = bcrypt.CompareHashAndPassword([]byte(*comp.Password), []byte(in.Password))
		if err != nil {
			if err == bcrypt.ErrMismatchedHashAndPassword {
				return nil, rpcerrors.CompetitionPasswordError
			}
			return nil, err
		}
		if attendance == nil {
			err = l.svcCtx.CompetitionAttendanceDao.CreateCompetitionAttendance(&models.CompetitionAttendance{
				CompetitionID:   in.CompetitionId,
				UserID:          in.UserId,
				PasswordVersion: &comp.PasswordVersion,
			})
		} else {
			err = l.svcCtx.CompetitionAttendanceDao.UpdateCompetitionAttendance(&models.CompetitionAttendance{
				CompetitionID:   in.CompetitionId,
				UserID:          in.UserId,
				PasswordVersion: &comp.PasswordVersion,
			})
		}
		if err != nil {
			return nil, err
		}
		return &competition.AttendCompetitionResp{}, nil
	} else {
		if attendance == nil {
			err = l.svcCtx.CompetitionAttendanceDao.CreateCompetitionAttendance(&models.CompetitionAttendance{
				CompetitionID:   in.CompetitionId,
				UserID:          in.UserId,
				PasswordVersion: nil,
			})
		} else {
			err = l.svcCtx.CompetitionAttendanceDao.UpdateCompetitionAttendance(&models.CompetitionAttendance{
				CompetitionID:   in.CompetitionId,
				UserID:          in.UserId,
				PasswordVersion: nil,
			})
		}
		if err != nil {
			return nil, err
		}
		return &competition.AttendCompetitionResp{}, nil
	}
}
