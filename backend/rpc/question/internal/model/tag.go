package model

import "time"

type Tag struct {
	ID        uint64    `gorm:"type:bigint unsigned;primaryKey;autoIncrement;column:id"`
	Name      string    `gorm:"type:varchar(64);unique;not null;column:name"`
	CreatedAt time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (Tag) TableName() string {
	return "tags"
}
