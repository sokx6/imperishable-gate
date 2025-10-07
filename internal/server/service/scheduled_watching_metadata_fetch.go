package service

import (
	"errors"
	"fmt"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"time"

	"gorm.io/gorm"
)

func ScheduledWatchingMetabaseFetch() {
	fmt.Println("Starting scheduled metadata fetch service...")
	ticker := time.NewTicker(3 * time.Hour)
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
			if err := database.DB.First(&user, link.UserID).Error; err != nil { // ✅ 用 First + .Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					fmt.Printf("User %d not found for link %s, skipping\n", link.UserID, link.Url)
				} else {
					fmt.Println("Error fetching user:", err)
				}
				continue
			}

			// ✅ 额外检查邮箱
			if user.Email == "" {
				fmt.Printf("User %d has no email, skipping notification\n", user.ID)
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
				fmt.Println("Error updating link metadata:", err)
			}
		}
	}
}
