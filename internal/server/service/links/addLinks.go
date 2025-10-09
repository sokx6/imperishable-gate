package links

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"time"

	"imperishable-gate/internal/server/service/common"

	"gorm.io/gorm"
)

// 并发控制：最多允许 100 个并发的爬虫 goroutine
var crawlerSemaphore = make(chan struct{}, 100)

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
			// 获取信号量，控制并发数
			crawlerSemaphore <- struct{}{}
			defer func() {
				<-crawlerSemaphore // 释放信号量
				if r := recover(); r != nil {
					logger.Error("Panic in metadata crawl for link %d: %v", linkId, r)
				}
			}()

			// 使用 timer 实现简单的超时控制
			done := make(chan bool, 1)
			go func() {
				title, desc, keywords, statusCode, _ := utils.CrawlMetadata(url)
				database.DB.Model(&model.Link{}).
					Where("id = ?", linkId).
					Updates(map[string]interface{}{
						"title":       title,
						"description": desc,
						"keywords":    keywords,
						"status_code": statusCode,
					})
				done <- true
			}()

			select {
			case <-done:
				logger.Debug("Metadata crawled successfully for link %d", linkId)
			case <-time.After(30 * time.Second):
				logger.Warning("Metadata crawl timeout for link %d (URL: %s)", linkId, url)
			}
		}()
		return nil
	}
	return common.ErrLinkAlreadyExists

}
