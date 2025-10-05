package service

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
)

func AddTagsByName(name string, userId uint, tags []string) error {
	tagList := utils.CreateTagList(tags, userId)

	var id uint
	if id = utils.NameToLinkId(name, userId); id == 0 {
		return ErrNameNotFound
	} else if err := database.DB.Model(&model.Link{ID: id}).Update("Tags", tagList).Error; err != nil {
		return ErrDatabase
	}
	return nil
}
