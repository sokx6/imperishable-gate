package tags

import (
	"errors"

	"gorm.io/gorm"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/service/common"
)

func DeleteTagsByLink(url string, userId uint, tags []string) error {
	var link model.Link
	if err := database.DB.Preload("Tags").Take(&link, "url = ? AND user_id = ?", url, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.ErrLinkNotFound
		}
		return common.ErrDatabase
	}
	existingTags := utils.ExtractTagNames(link.Tags)
	if !utils.ContainsAll(existingTags, tags) {
		return common.ErrInvalidRequest
	}

	var tagsToDelete []model.Tag
	if err := database.DB.Where("name IN ? AND user_id = ?", tags, userId).Find(&tagsToDelete).Error; err != nil {
		return common.ErrDatabase
	}

	if err := database.DB.Model(&link).Association("Tags").Delete(&tagsToDelete); err != nil {
		return common.ErrDatabase
	}

	return nil
}
