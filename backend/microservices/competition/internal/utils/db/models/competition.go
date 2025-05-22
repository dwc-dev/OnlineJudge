package models

import (
	"time"
)

type Competition struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement"`
	Name            string    `gorm:"type:varchar(100);not null"`
	Description     string    `gorm:"type:text;not null"`
	StartTime       time.Time `gorm:"not null"`
	EndTime         time.Time `gorm:"not null"`
	Questions       string    `gorm:"type:text;not null"`
	Type            string    `gorm:"type:enum('acm','oi');not null"`
	Password        *string   `gorm:"type:varchar(255)"`
	PasswordVersion uint      `gorm:"not null;default:0"`
	CreateAt        time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdateAt        time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (Competition) TableName() string {
	return "competitions"
}
