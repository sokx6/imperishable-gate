package database

import (
	"fmt"
	"imperishable-gate/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// 赋值给全局 DB
	DB = db

	// 局部 err 不影响全局
	if err := DB.AutoMigrate(&model.Link{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
