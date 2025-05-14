package svc

import (
	"backend/microservices/competition/internal/config"
	"backend/microservices/competition/internal/utils/db/dao"
	"backend/microservices/competition/internal/utils/db/models"
	"backend/microservices/judge/judgeclient"
	"backend/microservices/question/questionclient"
	"backend/microservices/user/userclient"
	"time"

	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config                   config.Config
	CompetitionDao           *dao.CompetitionDao
	CompetitionAttendanceDao *dao.CompetitionAttendanceDao
	CompetitionScoreDao      *dao.CompetitionScoreDao
	JudgeRpc                 judgeclient.Judge
	QuestionRpc              questionclient.Question
	UserRpc                  userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 获取底层连接池对象
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection pool")
	}
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)                  // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)                 // 最大打开连接数
	sqlDB.SetConnMaxLifetime(1 * time.Hour)    // 连接最大生命周期
	sqlDB.SetConnMaxIdleTime(30 * time.Minute) // 连接最大空闲时间
	// 自动迁移表结构
	db.AutoMigrate(&models.Competition{})
	db.AutoMigrate(&models.CompetitionAttendance{})
	db.AutoMigrate(&models.CompetitionScore{})
	return &ServiceContext{
		Config:                   c,
		CompetitionDao:           dao.NewCompetitionDao(db),
		CompetitionAttendanceDao: dao.NewCompetitionAttendanceDao(db),
		CompetitionScoreDao:      dao.NewCompetitionScoreDao(db),
		JudgeRpc:                 judgeclient.NewJudge(zrpc.MustNewClient(c.JudgeRpc)),
		QuestionRpc:              questionclient.NewQuestion(zrpc.MustNewClient(c.QuestionRpc)),
		UserRpc:                  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
