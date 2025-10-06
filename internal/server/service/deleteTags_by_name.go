package service

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
)

func DeleteTagsByName(name string, userId uint, tags []string) error {
	id := utils.GetLinkIDByName(name, userId)
	if id == 0 {
		return ErrNameNotFound
	}
	var link model.Link
	if err := database.DB.Preload("Tags").Take(&link, "id = ?", id).Error; err != nil {
		return ErrDatabase
	}
	existingTags := utils.ExtractTagNames(link.Tags)
	if !utils.ContainsAll(existingTags, tags) {
		return ErrInvalidRequest
	}

	var tagsToDelete []model.Tag
	if err := database.DB.Where("name IN ?", tags).Find(&tagsToDelete).Error; err != nil {
		return ErrDatabase
	}

	if err := database.DB.Model(&link).Association("Tags").Delete(&tagsToDelete); err != nil {
		return ErrDatabase
	}

	return nil
}
