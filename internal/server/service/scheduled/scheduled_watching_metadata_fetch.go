package scheduled

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"time"

	"gorm.io/gorm"
)

func ScheduledWatchingMetabaseFetch() {
	logger.Info("Starting scheduled watching metadata fetch service (interval: 3 hours)")
	ticker := time.NewTicker(3 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		var links []model.Link
		if err := database.DB.Where("watching = true").Find(&links).Error; err != nil {
			logger.Error("Error fetching watching links: %v", err)
			continue
		}
		logger.Info("Fetching metadata for %d watching links", len(links))
		for _, link := range links {
			logger.Debug("Processing link: ID=%d, URL=%s, UserID=%d", link.ID, link.Url, link.UserID)
			var user model.User
			if err := database.DB.First(&user, link.UserID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					logger.Warning("User %d not found for link %s, skipping", link.UserID, link.Url)
				} else {
					logger.Error("Error fetching user %d: %v", link.UserID, err)
				}
				continue
			}

			// 额外检查邮箱
			if user.Email == "" {
				logger.Warning("User %d has no email, skipping notification", user.ID)
				continue
			}
			title, desc, keywords, statusCode, _ := utils.CrawlMetadata(link.Url)
			if err := utils.CheckAndNotifyIfSiteChanged(
				link.Title,
				link.Description,
				link.Keywords,
				title,
				desc,
				keywords,
				user.Email,
				link.Url,
				link.StatusCode,
				statusCode); err != nil {

			}

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
