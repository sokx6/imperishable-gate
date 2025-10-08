package tags

import (
	"errors"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"

	"gorm.io/gorm"
	"imperishable-gate/internal/server/service/common"
)

func AddTagsByLink(url string, userId uint, tags []string) error {
	var link model.Link
	link.Url = url
	if err := database.DB.Where("url = ? AND user_id = ?", url, userId).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tagList := utils.CreateTagList(tags, userId)
			link.Tags = tagList
			link.UserID = userId
			if err := database.DB.Create(&link).Error; err != nil {
				return common.ErrDatabase
			}
			return nil
		} else {
			return common.ErrDatabase
		}
	}

	tagList := utils.CreateTagList(tags, userId)
	if len(tagList) == 0 {
		return nil
	}

	if err := database.DB.Model(&link).Association("Tags").Append(tagList); err != nil {
		return common.ErrDatabase
	}

	return nil
}
