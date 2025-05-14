package models

import (
	"time"
)

type CompetitionScore struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement"`
	CompetitionID uint64    `gorm:"not null"`
	UserID        uint64    `gorm:"not null"`
	ScoreDetails  string    `gorm:"type:text;not null" json:"score_details"` // 题目得分详情（JSON）
	JudgeIDs      string    `gorm:"type:text;not null" json:"judge_ids"`     // 评测 ID 列表（JSON 格式）
	TotalScore    float64   `gorm:"type:float;not null;default:0"`
	CreatedAt     time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (CompetitionScore) TableName() string {
	return "competition_scores"
}
