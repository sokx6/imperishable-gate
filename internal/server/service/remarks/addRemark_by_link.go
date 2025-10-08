package remarks

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"

	"gorm.io/gorm"
)

func AddRemarkByLink(url string, userId uint, remark string) error {
	var link model.Link
	result := database.DB.Where("url = ? AND user_id = ?", url, userId).First(&link)
	// 如果找不到对应的 Link，创建新的 Link 和关联的 Names
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		link = model.Link{
			UserID: userId,
			Url:    url,
			Remark: remark,
		}
		// 创建新的 Link
		if err := database.DB.Create(&link).Error; err != nil {
			return err
		}
	} else if result.Error != nil { // 其他数据库错误
		return result.Error
	}

	if err := database.DB.Model(&link).Update("Remark", remark).Error; err != nil {
		return err
	}

	return nil
}
