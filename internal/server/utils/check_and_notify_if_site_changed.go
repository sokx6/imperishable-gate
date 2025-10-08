package utils

import (
	"fmt"
	"imperishable-gate/internal/server/utils/email"
)

func CheckAndNotifyIfSiteChanged(
	oldTitle,
	oldDesc,
	oldKeywords,
	newTitle,
	newDesc,
	newKeywords,
	userEmail,
	changedUrl string,
	oldStatusCode,
	newStatusCode int) error {
	fmt.Println("Check change:", oldTitle, newTitle, oldDesc, newDesc, oldKeywords, newKeywords, oldStatusCode, newStatusCode)
	if oldTitle != newTitle ||
		oldDesc != newDesc ||
		oldKeywords != newKeywords ||
		oldStatusCode != newStatusCode {
		fmt.Println("Site changed, sending email...")
		if err := email.SendWebsiteChangeNotification(
			oldTitle,
			oldDesc,
			oldKeywords,
			newTitle,
			newDesc,
			newKeywords,
			userEmail,
			changedUrl,
			oldStatusCode,
			newStatusCode); err != nil {
			fmt.Println("Failed to send notification email:", err)
			return err
		}
	}
	return nil
}
