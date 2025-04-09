package model

import (
	"time"
)

type User struct {
	ID            uint64    `gorm:"type:bigint unsigned;primaryKey;autoIncrement"`
	UserName      string    `gorm:"type:varchar(50);unique;not null"`
	UserEmail     string    `gorm:"type:varchar(255);unique;not null"`
	UserPassword  string    `gorm:"type:varchar(255);not null"`
	UserAvatarURL string    `gorm:"type:varchar(1024);default:null"`
	UserProfile   string    `gorm:"type:text;default:null"`
	UserRole      string    `gorm:"type:enum('user', 'admin');default:'user';not null"`
	CreatedAt     time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null"`
}

func (User) TableName() string {
	return "users"
}
