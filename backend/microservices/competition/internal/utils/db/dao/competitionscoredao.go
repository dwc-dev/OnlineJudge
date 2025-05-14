package dao

import (
	"backend/microservices/competition/internal/utils/db/models"

	"gorm.io/gorm"
)

type CompetitionScoreDao struct {
	db *gorm.DB
}

func NewCompetitionScoreDao(db *gorm.DB) *CompetitionScoreDao {
	return &CompetitionScoreDao{db: db}
}

func (d *CompetitionScoreDao) CreateCompetitionScore(score *models.CompetitionScore) error {
	return d.db.Create(score).Error
}

func (d *CompetitionScoreDao) GetCompetitionScore(competitionId uint64, userId uint64) (*models.CompetitionScore, error) {
	var score models.CompetitionScore
	if err := d.db.Where("competition_id = ? AND user_id = ?", competitionId, userId).First(&score).Error; err != nil {
		return nil, err
	}
	return &score, nil
}

func (d *CompetitionScoreDao) SaveCompetitionScore(score *models.CompetitionScore) error {
	return d.db.Save(score).Error
}

func (d *CompetitionScoreDao) GetCompetitionScoreByCompetitionId(competitionId uint64) ([]*models.CompetitionScore, error) {
	var scores []*models.CompetitionScore
	if err := d.db.Where("competition_id = ?", competitionId).Order("total_score DESC").Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}
