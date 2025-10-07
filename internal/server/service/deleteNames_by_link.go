package service

import (
	"errors"

	"gorm.io/gorm"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
)

func DeleteNamesByLink(url string, userId uint, names []string) error {
	var link model.Link
	if err := database.DB.Preload("Names").Take(&link, "url = ? AND user_id = ?", url, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLinkNotFound
		}
		return ErrDatabase
	}
	existingNames := utils.ExtractNames(link.Names)
	if !utils.ContainsAll(existingNames, names) {
		return ErrInvalidRequest
	}

	if err := database.DB.Where("link_id = ? AND name IN ?", link.ID, names).Delete(&model.Name{}).Error; err != nil {
		return ErrDatabase
	}

	return nil
}
