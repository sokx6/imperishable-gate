package utils

import (
	"fmt"
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
		if err := sendEmail(
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
			fmt.Println("sendEmail 返回错误:", err)
			return err
		}
	}
	return nil
}
