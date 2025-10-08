package links

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"

	"gorm.io/gorm"
	"imperishable-gate/internal/server/service/common"
)

func AddLink(url string, userId uint) error {
	var link model.Link

	result := database.DB.Where("url = ? AND user_id = ?", url, userId).First(&link)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		link = model.Link{
			UserID: userId,
			Url:    url,
		}
		if err := database.DB.Create(&link).Error; err != nil {
			return common.ErrDatabase
		}
		linkId := link.ID
		go func() {
			defer func() {
				if r := recover(); r != nil {
					// 处理 panic，防止程序崩溃
				}
			}()
			// 爬取元数据并更新数据库
			title, desc, keywords, statusCode, _ := utils.CrawlMetadata(url)
			database.DB.Model(&model.Link{}).
				Where("id = ?", linkId).
				Updates(map[string]interface{}{
					"title":       title,
					"description": desc,
					"keywords":    keywords,
					"status_code": statusCode,
				})
		}()
		return nil
	}
	return common.ErrLinkAlreadyExists

}
