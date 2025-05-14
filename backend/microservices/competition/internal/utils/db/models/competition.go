package models

import (
	"time"
)

type Competition struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement"`
	Name            string     `gorm:"type:varchar(100);not null"`
	Description     string     `gorm:"type:text;not null"`
	StartTime       *time.Time `gorm:"type:timestamp"`
	EndTime         *time.Time `gorm:"type:timestamp"`
	Questions       string     `gorm:"type:text;not null"`
	Type            string     `gorm:"type:enum('acm','oi');not null"`
	Password        *string    `gorm:"type:varchar(255);default:null"`
	PasswordVersion uint       `gorm:"type:int unsigned;not null;default:0"`
	CreatedAt       time.Time  `gorm:"column:create_at;autoCreateTime"`
	UpdatedAt       time.Time  `gorm:"column:update_at;autoUpdateTime"`
}

func (Competition) TableName() string {
	return "competitions"
}
