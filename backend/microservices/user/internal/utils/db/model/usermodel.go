package model

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(50);unique;not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	AvatarURL string    `gorm:"type:varchar(1024);not null"`
	Profile   string    `gorm:"type:text;not null"`
	Role      string    `gorm:"type:enum('user','admin');default:'user';not null"`
	CreateAt  time.Time `gorm:"autoCreateTime"`
	UpdateAt  time.Time `gorm:"autoUpdateTime"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
