package service

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"

	"gorm.io/gorm"
)

func AddNames(url string, userId uint, names []string) error {
	var link model.Link
	link.Url = url

	if err := database.DB.Where("url = ? AND user_id = ?", url, userId).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 链接不存在，创建新链接并添加名称
			nameList := utils.CreateNameList(names, userId)
			link.Names = nameList
			link.UserID = userId
			if err := database.DB.Create(&link).Error; err != nil {
				return ErrDatabase
			}
			return nil
		} else {
			return ErrDatabase
		}
	}

	// 链接已存在，添加名称
	nameList := utils.CreateNameList(names, userId)
	if len(nameList) == 0 {
		return ErrInvalidRequest
	}

	if err := database.DB.Model(&link).Association("Names").Append(nameList); err != nil {
		return ErrNameAlreadyExists
	}

	return nil
}
