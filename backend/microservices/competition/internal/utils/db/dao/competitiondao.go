package dao

import (
	"backend/microservices/competition/internal/utils/db/models"

	"gorm.io/gorm"
)

type CompetitionDao struct {
	db *gorm.DB
}

func NewCompetitionDao(db *gorm.DB) *CompetitionDao {
	return &CompetitionDao{db: db}
}

func (d *CompetitionDao) CreateCompetition(competition *models.Competition) error {
	return d.db.Create(competition).Error
}

func (d *CompetitionDao) GetCompetitionById(id uint64) (*models.Competition, error) {
	var competition models.Competition
	if err := d.db.Where("id = ?", id).First(&competition).Error; err != nil {
		return nil, err
	}
	return &competition, nil
}

func (d *CompetitionDao) GetPasswordVersion(id uint64) (uint64, error) {
	var passwordVersion uint64
	if err := d.db.Model(&models.Competition{}).Where("id = ?", id).Select("password_version").First(&passwordVersion).Error; err != nil {
		return 0, err
	}
	return passwordVersion, nil
}

func (d *CompetitionDao) GetCompetitionList(page, pageSize int64, filter map[string]string, col []string) ([]*models.Competition, int64, error) {
	var competitions []*models.Competition
	query := d.db.Model(&models.Competition{})

	if len(col) > 0 {
		query = query.Select(col)
	}

	if name, ok := filter["name"]; ok && name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if ctype, ok := filter["type"]; ok && ctype != "" {
		query = query.Where("type = ?", ctype)
	}

	if description, ok := filter["description"]; ok && description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Order("start_time DESC").Find(&competitions).Error; err != nil {
		return nil, 0, err
	}

	return competitions, total, nil
}

func (d *CompetitionDao) UpdateCompetition(competition *models.Competition) error {
	err := d.db.Model(&models.Competition{}).Where("id = ?", competition.ID).Updates(competition).Error
	if err != nil {
		return err
	}
	if competition.Password == nil {
		return d.db.Model(&models.Competition{}).Where("id = ?", competition.ID).Update("password", nil).Error
	}
	return nil
}

func (d *CompetitionDao) DeleteCompetition(id uint64) error {
	return d.db.Delete(&models.Competition{}, id).Error
}
