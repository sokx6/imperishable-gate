package service

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"

	"gorm.io/gorm"
)

func AddLink(url string, userId uint) error {
	var link model.Link

	result := database.DB.Where("url = ? AND user_id = ?", url, userId).First(&link)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		title, desc, keywords := utils.CrawlMetadata(url)
		link = model.Link{
			UserID:      userId,
			Url:         url,
			Title:       title,
			Description: desc,
			Keywords:    keywords,
		}
		if err := database.DB.Create(&link).Error; err != nil {
			return ErrDatabase
		}
		return nil
	}
	return ErrLinkAlreadyExists

}
