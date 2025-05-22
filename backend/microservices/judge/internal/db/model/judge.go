package model

import (
	"time"
)

type Judge struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement"`
	QuestionID uint64    `gorm:"not null"`
	UserID     uint64    `gorm:"not null"`
	Language   string    `gorm:"type:enum('c','cpp','java','python','golang','rust');not null"`
	Code       string    `gorm:"type:text;not null"`
	ExecResult string    `gorm:"type:text;not null"`
	Accepted   bool      `gorm:"not null"`
	JudgeType  string    `gorm:"type:enum('normal','competition');not null"`
	CreateAt   time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (Judge) TableName() string {
	return "judges"
}
