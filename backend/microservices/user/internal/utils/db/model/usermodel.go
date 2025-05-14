package model

import (
	"time"
)

type User struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement;column:id;comment:用户ID"`
	UserName      string    `gorm:"type:varchar(50);unique;not null;column:user_name;comment:用户昵称"`
	UserEmail     string    `gorm:"type:varchar(255);unique;not null;column:user_email;comment:用户邮箱"`
	UserPassword  string    `gorm:"type:varchar(255);not null;column:user_password;comment:密码"`
	UserAvatarURL string    `gorm:"type:varchar(1024);not null;column:user_avatar_url;comment:用户头像URL"`
	UserProfile   *string   `gorm:"type:text;column:user_profile;comment:用户简介"`
	UserRole      string    `gorm:"type:enum('user','admin');default:user;not null;column:user_role;comment:用户角色：user/admin"`
	SummitNum     uint64    `gorm:"type:bigint unsigned;not null;default:0;column:summit_num;comment:提交数"`
	AcceptedNum   uint64    `gorm:"type:bigint unsigned;not null;default:0;column:accepted_num;comment:通过数"`
	QutsionTryNum uint64    `gorm:"type:bigint unsigned;not null;default:0;column:qutsion_try_num;comment:尝试的题目数"`
	CreateAt      time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;column:create_at;comment:创建时间"`
	UpdateAt      time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:update_at;comment:更新时间"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
