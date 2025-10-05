package service

import (
	"fmt"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"time"
)

func ScheduledMetabaseFetch() {
	fmt.Println("Starting scheduled metadata fetch service...")
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		var links []model.Link
		if err := database.DB.Find(&links).Error; err != nil {
			fmt.Println("Error fetching links:", err)
			continue
		}
		for _, link := range links {
			fmt.Println("Fetching metadata for URL:", link.Url)
			title, desc, keywords := utils.CrawlMetadata(link.Url)
			link.Title = title
			link.Description = desc
			link.Keywords = keywords
			if err := database.DB.Save(&link).Error; err != nil {
				fmt.Println("Error updating link metadata:", err)
			}
		}
	}
}
