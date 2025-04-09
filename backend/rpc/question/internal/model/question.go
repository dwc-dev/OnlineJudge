package model

import "time"

type Question struct {
	ID          uint64 `gorm:"type:bigint unsigned;primaryKey;autoIncrement;column:id"`
	Title       string `gorm:"type:varchar(512);not null;column:title"`
	Content     string `gorm:"type:text;not null;column:content"`
	Answer      string `gorm:"type:text;column:answer"`
	Tags        string `gorm:"type:text;column:tags"`
	Difficulty  string `gorm:"type:enum('easy','medium','hard');not null;column:difficulty"`
	SubmitNum   uint64 `gorm:"type:bigint unsigned;not null;default:0;column:submit_num"`
	AcceptedNum uint64 `gorm:"type:bigint unsigned;not null;default:0;column:accepted_num"`
	JudgeCase   string `gorm:"type:text;not null;column:judge_case"`
	JudgeConfig string `gorm:"type:text;column:judge_config"`
	// ThumbNum     uint64    `gorm:"type:bigint unsigned;not null;default:0;column:thumb_num"`
	// FavourNum    uint64    `gorm:"type:bigint unsigned;not null;default:0;column:favour_num"`
	CreateUserID uint64    `gorm:"type:bigint unsigned;not null;column:create_user_id"`
	CreatedAt    time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (Question) TableName() string {
	return "questions"
}
