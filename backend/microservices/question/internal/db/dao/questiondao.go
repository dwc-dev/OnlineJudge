package dao

import (
	"backend/microservices/question/internal/db/model"

	"gorm.io/gorm"
)

type QuestionDao struct {
	db *gorm.DB
}

func NewQuestionDao(db *gorm.DB) *QuestionDao {
	return &QuestionDao{db: db}
}

func (d *QuestionDao) AddQuestion(question *model.Question) error {
	return d.db.Create(question).Error
}

func (d *QuestionDao) DeleteQuestion(id uint64) error {
	return d.db.Delete(&model.Question{}, id).Error
}

func (d *QuestionDao) UpdateQuestion(question *model.Question) error {
	return d.db.Model(&model.Question{}).Where("id = ?", question.ID).Updates(question).Error
}

func (d *QuestionDao) GetQuestionInfo(id uint64, filter map[string]string, col []string) (*model.Question, error) {
	var question model.Question
	query := d.db.Model(&model.Question{})

	if len(col) > 0 {
		query = query.Select(col)
	}

	if visible_scope, ok := filter["visible_scope"]; ok && visible_scope != "" {
		query = query.Where("visible_scope = ?", visible_scope)
	}

	if err := query.Where("id = ?", id).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (d *QuestionDao) GetQuestionList(page, pageSize int64, filter map[string]string, col []string) ([]*model.Question, int64, error) {
	var questions []*model.Question
	query := d.db.Model(&model.Question{})

	if len(col) > 0 {
		query = query.Select(col)
	}

	if title, ok := filter["title"]; ok && title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	if difficulty, ok := filter["difficulty"]; ok && difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}

	if tag, ok := filter["tag"]; ok && tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	if visible_scope, ok := filter["visible_scope"]; ok && visible_scope != "" {
		query = query.Where("visible_scope = ?", visible_scope)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&questions).Error; err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

func (d *QuestionDao) GetQuestionJudgeCase(id uint64) (*model.Question, error) {
	var question model.Question
	if err := d.db.Select("judge_case").Where("id = ?", id).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (d *QuestionDao) GetQuestionJudgeConfig(id uint64) (*model.Question, error) {
	var question model.Question
	if err := d.db.Select("judge_config").Where("id = ?", id).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}
