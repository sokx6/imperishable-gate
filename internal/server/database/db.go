package database

import (
	"fmt"
	"imperishable-gate/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	// 局部 err 不影响全局
	var err error

	// 连接数据库
	// dsn是指定的数据库连接字符串
	// 例如 "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// gorm.Open 返回一个 *gorm.DB 实例和一个 error
	// gorm.Config 可以用来配置 GORM 的行为，这里使用默认配置
	// 连接成功后，db 就是一个可以用来操作数据库的对象
	// 如果连接失败，err 会包含错误信息
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// 赋值给全局 DB
	DB = db

	// 自动迁移数据库
	// AutoMigrate 会根据 model.Link 结构体的定义
	// 自动创建或更新数据库中的表结构
	if err := DB.AutoMigrate(&model.Tag{}, &model.Link{}, &model.Name{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
