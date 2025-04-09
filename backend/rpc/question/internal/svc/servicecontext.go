package svc

import (
	"backend/rpc/question/internal/config"
	"backend/rpc/question/internal/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB // 添加 gorm.DB
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
	db.AutoMigrate(&model.Question{})

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
