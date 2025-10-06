package service

import (
	"fmt"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"time"
)

func ScheduledWatchingMetabaseFetch() {
	fmt.Println("Starting scheduled metadata fetch service...")
	ticker := time.NewTicker(6 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var links []model.Link
		if err := database.DB.Where("watching = true").Find(&links).Error; err != nil {
			fmt.Println("Error fetching links:", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
			fmt.Println("Fetching metadata for URL:", link.Url)
			var user model.User
			if err := database.DB.Find(&user, link.UserID); err != nil {
			}
			title, desc, keywords, statusCode := utils.CrawlMetadata(link.Url)
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
				fmt.Println("Error updating link metadata:", err)
			}
		}
	}
}
