package database

import (
	"fmt"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/utils/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	// 局部 err 不影响全局
	var err error

	logger.Info("Initializing database connection...")

	// 连接数据库
	// dsn是指定的数据库连接字符串
	// 例如 "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// gorm.Open 返回一个 *gorm.DB 实例和一个 error
	// gorm.Config 可以用来配置 GORM 的行为，这里使用默认配置
	// 连接成功后，db 就是一个可以用来操作数据库的对象
	// 如果连接失败，err 会包含错误信息
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// 赋值给全局 DB
	DB = db
	logger.Success("Database connected successfully")

	// 自动迁移数据库
	// AutoMigrate 会根据 model.Link 结构体的定义
	// 自动创建或更新数据库中的表结构
	logger.Info("Running database migrations...")
	if err := DB.AutoMigrate(&model.Tag{},
		&model.Link{},
		&model.Name{},
		&model.LinkTag{},
		&model.User{},
		&model.RefreshToken{},
		&model.EmailVerification{}); err != nil {
		logger.Error("Failed to migrate database: %v", err)
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	logger.Success("Database migrations completed successfully")
	return nil
}
