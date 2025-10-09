package email

import (
	"crypto/tls"
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

	// 根据端口选择发送方式
	if config.SMTPPort == "465" {
		// 使用 SSL/TLS 加密连接（端口 465）
		err = sendMailWithTLS(config, auth, []string{to}, []byte(message))
	} else {
		// 使用 STARTTLS（端口 587）或明文（端口 25，不推荐）
		err = smtp.SendMail(config.GetSMTPAddress(), auth, config.From, []string{to}, []byte(message))
	}

	if err != nil {
		logger.Error("Failed to send email to %s: %v", to, err)
		return err
	}
	logger.Success("Email sent successfully to: %s", to)
	return nil
}

// sendMailWithTLS 使用 TLS 加密连接发送邮件（用于端口 465）
func sendMailWithTLS(config *Config, auth smtp.Auth, to []string, msg []byte) error {
	// 创建 TLS 配置
	tlsConfig := &tls.Config{
		ServerName: config.SMTPHost,
	}

	// 建立 TLS 连接
	conn, err := tls.Dial("tcp", config.GetSMTPAddress(), tlsConfig)
	if err != nil {
		return fmt.Errorf("TLS connection failed: %v", err)
	}
	defer conn.Close()

	// 创建 SMTP 客户端
	client, err := smtp.NewClient(conn, config.SMTPHost)
	if err != nil {
		return fmt.Errorf("SMTP client creation failed: %v", err)
	}
	defer client.Close()

	// 认证
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %v", err)
	}

	// 设置发件人
	if err = client.Mail(config.From); err != nil {
		return fmt.Errorf("setting sender failed: %v", err)
	}

	// 设置收件人
	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			return fmt.Errorf("setting recipient failed: %v", err)
		}
	}

	// 发送邮件内容
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("data command failed: %v", err)
	}

	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("writing message failed: %v", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("closing data writer failed: %v", err)
	}

	return client.Quit()
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
