package database

import (
	"fmt"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/utils/logger"
	"os"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	// 局部 err 不影响全局
	var err error

	logger.Info("Initializing database connection...")

	// 获取数据库类型，默认为 sqlite
	dbType := strings.ToLower(os.Getenv("DB_TYPE"))
	if dbType == "" {
		dbType = "sqlite"
	}

	// 根据数据库类型选择对应的驱动
	var dialector gorm.Dialector
	switch dbType {
	case "mysql":
		dialector = mysql.Open(dsn)
		logger.Info("Using MySQL database")
	case "postgres", "postgresql":
		dialector = postgres.Open(dsn)
		logger.Info("Using PostgreSQL database")
	case "sqlite":
		// 使用纯 Go 实现的 SQLite 驱动（无需 CGO）
		dialector = sqlite.Open(dsn)
		logger.Info("Using SQLite database (pure Go driver)")
	default:
		return fmt.Errorf("unsupported database type: %s (supported: mysql, postgres, sqlite)", dbType)
	}

	// 连接数据库
	// dsn 是指定的数据库连接字符串
	// SQLite: 例如 "gate.db" 或 "/path/to/gate.db" 或 "file:gate.db?cache=shared&mode=rwc"
	// MySQL: 例如 "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// PostgreSQL: 例如 "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	// 获取日志级别配置
	logLevel := getGormLogLevel()

	// 创建自定义 GORM logger
	gormLog := logger.NewGormLogger(logLevel, 200*time.Millisecond)

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLog,
	})
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

// getGormLogLevel 根据环境变量获取 GORM 日志级别
func getGormLogLevel() gormLogger.LogLevel {
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	switch logLevel {
	case "debug":
		return gormLogger.Info // GORM Info 级别会记录所有 SQL
	case "info":
		return gormLogger.Warn // 只记录慢查询和错误
	case "warn", "warning":
		return gormLogger.Warn
	case "error":
		return gormLogger.Error
	case "silent":
		return gormLogger.Silent
	default:
		// 默认为 Warn 级别，只记录慢查询和错误
		return gormLogger.Warn
	}
}
