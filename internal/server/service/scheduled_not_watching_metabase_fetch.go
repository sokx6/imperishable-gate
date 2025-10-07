package service

import (
	"fmt"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"time"
)

func ScheduledNotWatchingMetabaseFetch() {
	fmt.Println("Starting scheduled metadata fetch service...")
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		var links []model.Link
		if err := database.DB.Where("watching = false").Find(&links).Error; err != nil {
			fmt.Println("Error fetching links:", err)
			continue
		}
		for i := range links {
			link := &links[i]
			fmt.Println("Fetching metadata for URL:", link.Url)
			title, desc, keywords, statusCode, _ := utils.CrawlMetadata(link.Url)
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
