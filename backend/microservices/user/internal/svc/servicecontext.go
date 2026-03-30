package svc

import (
	"backend/microservices/user/internal/config"
	"backend/microservices/user/internal/utils/db/dao"
	"backend/microservices/user/internal/utils/db/model"
	"backend/microservices/user/internal/utils/mc"
	"backend/microservices/user/internal/utils/redis"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	UserDao     *dao.UserDao
	MinioClient *mc.MinioClient
	RedisClient *redis.RedisClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + fmt.Sprintf("%s", c.Mysql.DataSource))
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
	db.AutoMigrate(&model.User{})

	// 创建minio客户端
	minioClient, err := mc.NewMinioClient(c.Minio)
	if err != nil {
		panic("failed to create minio client")
	}
	redisClient := redis.NewRedisClient(c.MyRedis)
	return &ServiceContext{
		Config:      c,
		UserDao:     dao.NewUserDao(db),
		MinioClient: minioClient,
		RedisClient: redisClient,
	}
}
