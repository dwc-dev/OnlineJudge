package model

import "time"

type QuestionTagRelation struct {
	QuestionID uint64    `gorm:"type:bigint unsigned;primaryKey;column:question_id"`
	TagID      uint64    `gorm:"type:bigint unsigned;primaryKey;column:tag_id"`
	CreatedAt  time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (QuestionTagRelation) TableName() string {
	return "question_tag_relations"
}
