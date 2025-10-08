package tags

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/service/common"
)

func DeleteTagsByName(name string, userId uint, tags []string) error {
	id := utils.GetLinkIDByName(name, userId)
	if id == 0 {
		return common.ErrNameNotFound
	}
	var link model.Link
	if err := database.DB.Preload("Tags").Take(&link, "id = ?", id).Error; err != nil {
		return common.ErrDatabase
	}
	existingTags := utils.ExtractTagNames(link.Tags)
	if !utils.ContainsAll(existingTags, tags) {
		return common.ErrInvalidRequest
	}

	var tagsToDelete []model.Tag
	if err := database.DB.Where("name IN ?", tags).Find(&tagsToDelete).Error; err != nil {
		return common.ErrDatabase
	}

	if err := database.DB.Model(&link).Association("Tags").Delete(&tagsToDelete); err != nil {
		return common.ErrDatabase
	}

	return nil
}
