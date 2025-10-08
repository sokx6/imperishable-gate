package email

import (
	"fmt"
	"imperishable-gate/internal/server/utils/logger"
	"mime"
	"net/smtp"
)

// Send 通用邮件发送函数
func SendEmail(to, subject, htmlBody string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// 设置邮件头部
	header := make(map[string]string)
	header["From"] = config.From
	header["To"] = to
	header["Subject"] = mime.QEncoding.Encode("UTF-8", subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + htmlBody

	// 认证信息
	auth := smtp.PlainAuth("", config.From, config.Password, config.SMTPHost)

	// 发送邮件
	logger.Info("Sending email to: %s, subject: %s", to, subject)
	err = smtp.SendMail(config.GetSMTPAddress(), auth, config.From, []string{to}, []byte(message))
	if err != nil {
		logger.Error("Failed to send email to %s: %v", to, err)
		return err
	}
	logger.Success("Email sent successfully to: %s", to)
	return nil
}

// SendVerificationEmail 发送验证邮件
func SendVerificationEmail(to, code string) error {
	subject := " Imperishable Gate - 邮箱验证码"
	htmlBody := GetVerificationEmailTemplate(code)
	return SendEmail(to, subject, htmlBody)
}

// SendWebsiteChangeNotification 发送网站变更通知邮件
func SendWebsiteChangeNotification(
	oldTitle, oldDesc, oldKeywords,
	newTitle, newDesc, newKeyword,
	userEmail, changedUrl string,
	oldStatusCode, newStatusCode int) error {

	// 邮件主题：动态提示
	var subject string
	if oldStatusCode != newStatusCode {
		subject = "网站状态变更提醒: " + newTitle
	} else {
		subject = "您关注的页面有更新啦: " + newTitle
	}

	htmlBody := GetWebsiteChangeEmailTemplate(
		oldTitle, oldDesc, oldKeywords,
		newTitle, newDesc, newKeyword,
		changedUrl,
		oldStatusCode, newStatusCode,
	)

	logger.Info("Preparing to send website change notification to: %s", userEmail)
	return SendEmail(userEmail, subject, htmlBody)
}
