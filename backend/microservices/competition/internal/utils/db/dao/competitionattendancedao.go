package dao

import (
	"backend/microservices/competition/internal/utils/db/models"

	"gorm.io/gorm"
)

type CompetitionAttendanceDao struct {
	db *gorm.DB
}

func NewCompetitionAttendanceDao(db *gorm.DB) *CompetitionAttendanceDao {
	return &CompetitionAttendanceDao{db: db}
}

func (d *CompetitionAttendanceDao) GetCompetitionAttendance(competitionId uint64, userId uint64) (*models.CompetitionAttendance, error) {
	var attendance models.CompetitionAttendance
	if err := d.db.Where("competition_id = ? AND user_id = ?", competitionId, userId).First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (d *CompetitionAttendanceDao) CreateCompetitionAttendance(attendance *models.CompetitionAttendance) error {
	return d.db.Create(attendance).Error
}

func (d *CompetitionAttendanceDao) UpdateCompetitionAttendance(attendance *models.CompetitionAttendance) error {
	return d.db.Model(&models.CompetitionAttendance{}).Where("competition_id = ? AND user_id = ?", attendance.CompetitionID, attendance.UserID).Updates(attendance).Error
}

func (d *CompetitionAttendanceDao) GetCompetitionAttendancesByUserId(userId uint64) ([]*models.CompetitionAttendance, error) {
	var attendances []*models.CompetitionAttendance
	if err := d.db.Where("user_id = ?", userId).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}
