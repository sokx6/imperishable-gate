package scheduled

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"time"
)

func ScheduledNotWatchingMetabaseFetch() {
	logger.Info("Starting scheduled non-watching metadata fetch service (interval: 24 hours)")
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		var links []model.Link
		if err := database.DB.Where("watching = false").Find(&links).Error; err != nil {
			logger.Error("Error fetching non-watching links: %v", err)
			continue
		}
		logger.Info("Fetching metadata for %d non-watching links", len(links))
		for i := range links {
			link := &links[i]
			logger.Debug("Fetching metadata for URL: %s", link.Url)
			title, desc, keywords, statusCode, _ := utils.CrawlMetadata(link.Url)
			link.Title = title
			link.Description = desc
			link.Keywords = keywords
			link.StatusCode = statusCode
			if err := database.DB.Save(&link).Error; err != nil {
				logger.Error("Error updating link metadata for URL %s: %v", link.Url, err)
			} else {
				logger.Success("Updated metadata for link: %s", link.Url)
			}
		}
	}
}
