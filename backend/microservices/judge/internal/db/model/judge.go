package model

import (
	"time"
)

type Judge struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	QuestionID uint64    `gorm:"not null;column:question_id" json:"question_id"`
	UserID     uint64    `gorm:"not null;column:user_id" json:"user_id"`
	Language   string    `gorm:"type:enum('c','cpp','java','python','golang','rust');column:language" json:"language"`
	Code       string    `gorm:"type:text;not null;column:code" json:"code"`
	ExecResult string    `gorm:"type:text;not null;column:exec_result" json:"exec_result"`
	Accepted   bool      `gorm:"not null;column:accepted" json:"accepted"`
	JudgeType  string    `gorm:"type:enum('normal','competition');column:judge_type" json:"judge_type"`
	CreatedAt  time.Time `gorm:"column:create_at;autoCreateTime" json:"create_at"`
	UpdatedAt  time.Time `gorm:"column:update_at;autoUpdateTime" json:"update_at"`
}

func (Judge) TableName() string {
	return "judges"
}
