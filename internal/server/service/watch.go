package service

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"

	"gorm.io/gorm"
)

func Watch(url string, userId uint, watch bool) error {
	var link model.Link
	link.Url = url

	if err := database.DB.Where("url = ? AND user_id = ?", url, userId).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLinkNotFound
		} else {
			return ErrDatabase
		}
	}

	link.Watching = watch

	if err := database.DB.Save(&link).Error; err != nil {
		return ErrDatabase
	}

	return nil
}
