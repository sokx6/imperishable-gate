package utils

import (
	"imperishable-gate/internal/server/utils/email"
	"imperishable-gate/internal/server/utils/logger"
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
	logger.Debug("Checking site changes - URL: %s, Old Status: %d, New Status: %d", changedUrl, oldStatusCode, newStatusCode)
	if oldTitle != newTitle ||
		oldDesc != newDesc ||
		oldKeywords != newKeywords ||
		oldStatusCode != newStatusCode {
		logger.Info("Site changed detected for %s, sending notification to %s", changedUrl, userEmail)
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
			logger.Error("Failed to send notification email to %s: %v", userEmail, err)
			return err
		}
		logger.Success("Notification email sent successfully to %s", userEmail)
	}
	return nil
}
