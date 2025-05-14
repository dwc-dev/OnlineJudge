package model

import (
	"time"
)

type Question struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Title        string    `gorm:"type:varchar(512);not null;column:title" json:"title"`
	Content      string    `gorm:"type:text;not null;column:content" json:"content"`
	Tags         string    `gorm:"type:text;not null;column:tags" json:"tags"`
	Difficulty   string    `gorm:"type:enum('easy','medium','hard');not null;column:difficulty" json:"difficulty"`
	SubmitNum    uint64    `gorm:"type:bigint unsigned;not null;default:0;column:submit_num" json:"submit_num"`
	AcceptedNum  uint64    `gorm:"type:bigint unsigned;not null;default:0;column:accepted_num" json:"accepted_num"`
	JudgeCase    string    `gorm:"type:text;not null;column:judge_case" json:"judge_case"`
	JudgeConfig  string    `gorm:"type:text;not null;column:judge_config" json:"judge_config"`
	VisibleScope string    `gorm:"type:enum('public','competition_only');not null;default:'public';column:visible_scope" json:"visible_scope"`
	CreateAt     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;column:create_at" json:"create_at"`
	UpdateAt     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:update_at" json:"update_at"`
}

// TableName 设置表名
func (Question) TableName() string {
	return "questions"
}
