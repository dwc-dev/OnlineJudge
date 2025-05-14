package models

import (
	"time"
)

type CompetitionAttendance struct {
	CompetitionID   uint64    `gorm:"primaryKey"`
	UserID          uint64    `gorm:"primaryKey"`
	PasswordVersion *uint     `gorm:"type:int unsigned" json:"password_version,omitempty"`
	CreateAt        time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;column:create_at;comment:创建时间"`
	UpdateAt        time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:update_at;comment:更新时间"`
}

func (CompetitionAttendance) TableName() string {
	return "competition_attendances"
}
