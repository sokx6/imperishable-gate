package links

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"

	"gorm.io/gorm"
	"imperishable-gate/internal/server/service/common"
)

func Watch(url string, userId uint, watch bool) error {
	var link model.Link
	link.Url = url

	if err := database.DB.Where("url = ? AND user_id = ?", url, userId).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.ErrLinkNotFound
		} else {
			return common.ErrDatabase
		}
	}

	link.Watching = watch

	if err := database.DB.Save(&link).Error; err != nil {
		return common.ErrDatabase
	}

	return nil
}
