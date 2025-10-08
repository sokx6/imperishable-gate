package utils

import (
	"imperishable-gate/internal/server/utils/email"
)

func GetVerificationEmailTemplate(code string) string {
	return email.GetVerificationEmailTemplate(code)
}

func SendEmail(to, subject, body string) error {
	return email.SendEmail(to, subject, body)
}
