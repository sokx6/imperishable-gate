package service

import (
	"errors"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"

	"gorm.io/gorm"
)

func AddTagsByLink(url string, tags []string) error {
	var link model.Link
	link.Url = url
	if err := database.DB.Where("url = ?", url).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tagList := utils.CreateTagList(tags)
			link.Tags = tagList
			if err := database.DB.Create(&link).Error; err != nil {
				return ErrDatabase
			}
			return nil
		} else {
			return ErrDatabase
		}
	}

	tagList := utils.CreateTagList(tags)
	if len(tagList) == 0 {
		return nil
	}

	if err := database.DB.Model(&link).Association("Tags").Append(tagList); err != nil {
		return ErrDatabase
	}

	return nil
}
