package service

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
)

func AddTagsByName(name string, tags []string) error {
	tagList := utils.CreateTagList(tags)

	var Name model.Name
	if err := database.DB.Where("name = ?", name).Take(&Name).Error; err != nil {
		return ErrNameNotFound
	}
	if err := database.DB.Model(&model.Link{ID: Name.LinkID}).Update("Tags", tagList).Error; err != nil {
		return ErrDatabase
	}
	return nil
}
