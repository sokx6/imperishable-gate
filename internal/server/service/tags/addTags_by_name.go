package tags

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/service/common"
)

func AddTagsByName(name string, userId uint, tags []string) error {
	tagList := utils.CreateTagList(tags, userId)

	var id uint
	if id = utils.GetLinkIDByName(name, userId); id == 0 {
		return common.ErrNameNotFound
	} else if err := database.DB.Model(&model.Link{ID: id}).Update("Tags", tagList).Error; err != nil {
		return common.ErrDatabase
	}
	return nil
}
