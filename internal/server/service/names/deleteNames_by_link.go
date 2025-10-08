package names

import (
	"errors"

	"gorm.io/gorm"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/service/common"
)

func DeleteNamesByLink(url string, userId uint, names []string) error {
	var link model.Link
	if err := database.DB.Preload("Names").Take(&link, "url = ? AND user_id = ?", url, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.ErrLinkNotFound
		}
		return common.ErrDatabase
	}
	existingNames := utils.ExtractNames(link.Names)
	if !utils.ContainsAll(existingNames, names) {
		return common.ErrInvalidRequest
	}

	if err := database.DB.Where("link_id = ? AND name IN ?", link.ID, names).Delete(&model.Name{}).Error; err != nil {
		return common.ErrDatabase
	}

	return nil
}
