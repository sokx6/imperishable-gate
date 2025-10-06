package utils

import (
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendEmail(
	oldTitle,
	oldDesc,
	oldKeywords,
	newTitle,
	newDesc,
	newKeyword,
	userEmail,
	changedUrl string,
	oldStatusCode,
	newStatusCode int) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	fmt.Println("starting")

	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	// 使用传入的真实用户邮箱作为收件人
	to := []string{userEmail}

	// 🔔 邮件主题：带 emoji 和动态提示
	var subject string
	if oldStatusCode != newStatusCode {
		subject = "⚠️ 网站状态变更提醒: " + newTitle
	} else {
		subject = "🎉 您关注的页面有更新啦: " + newTitle
	}

	// 🎨 HTML 邮件正文（更美观）
	body := `
	<html>
	<body style="font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; color: #333; line-height: 1.6; max-width: 700px; margin: auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 12px; background-color: #f9f9fb;">
		<div style="text-align: center; margin-bottom: 20px;">
			<h2 style="color: #4a90e2;">🚀 网站监控小助手</h2>
			<p style="color: #666; font-size: 14px;">第一时间为您推送重要变化 💬</p>
			<hr style="border: 1px dashed #ccc;" />
		</div>

		<h3 style="color: #2c3e50;">📌 页面已更新</h3>
		<p><strong>📍 网址:</strong> <a href="` + changedUrl + `" style="color: #4a90e2; text-decoration: none;">访问此页</a></p>

		<table width="100%" style="background-color: #f1f8ff; padding: 15px; border-radius: 8px; border: 1px solid #d0ebff; margin: 10px 0;">
			<tr>
				<td style="padding: 8px;"><strong>📌 项目</strong></td>
				<td style="padding: 8px;"><strong>旧内容</strong></td>
				<td style="padding: 8px;"><strong>新内容 🎉</strong></td>
			</tr>
			<tr>
				<td style="padding: 8px;">📄 标题</td>
				<td style="padding: 8px; color: #999;"><i>` + htmlEscape(oldTitle) + `</i></td>
				<td style="padding: 8px; color: #27ae60; font-weight: bold;">` + htmlEscape(newTitle) + `</td>
			</tr>
			<tr>
				<td style="padding: 8px;">📝 描述</td>
				<td style="padding: 8px; color: #999; white-space: pre-wrap;"><i>` + htmlEscape(truncateText(oldDesc, 100)) + `</i></td>
				<td style="padding: 8px; color: #27ae60; white-space: pre-wrap;">` + htmlEscape(truncateText(newDesc, 100)) + `</td>
			</tr>
			<tr>
				<td style="padding: 8px;">🔖 关键词</td>
				<td style="padding: 8px; color: #999;"><i>` + htmlEscape(truncateText(oldKeywords, 80)) + `</i></td>
				<td style="padding: 8px; color: #27ae60;">` + htmlEscape(truncateText(newKeyword, 80)) + `</td>
			</tr>
			<tr>
				<td style="padding: 8px;">🔧 状态码</td>
				<td style="padding: 8px; color: ` + statusCodeColor(oldStatusCode) + `;"><b>` + fmt.Sprintf("%d", oldStatusCode) + `</b></td>
				<td style="padding: 8px; color: ` + statusCodeColor(newStatusCode) + `;"><b>` + fmt.Sprintf("%d", newStatusCode) + `</b></td>
			</tr>
		</table>

		<div style="margin-top: 20px; text-align: center; font-size: 14px; color: #777;">
			<p>🔔 这是自动通知，请勿回复本邮件。<br/>
			如有疑问，欢迎联系支持团队 ❤️</p>
			<small>Powered by Your Site Monitor ✨</small>
		</div>
	</body>
	</html>
	`

	// 设置邮件头部（使用 MIME 多部分格式支持 HTML）
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to[0]
	header["Subject"] = mime.QEncoding.Encode("UTF-8", subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 认证信息
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 发送邮件
	fmt.Println("准备发送邮件，SMTP配置：", smtpHost, smtpPort, from, userEmail)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println("邮件发送失败:", err)
		return err
	}
	log.Println("✅ 邮件已成功发送至:", userEmail)
	return nil
}

// 工具函数：转义 HTML 特殊字符，防止 XSS 注入或格式错乱
func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

// 工具函数：截断过长的文本，避免表格太宽
func truncateText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}
	return text[:maxLength] + "..."
}

// 工具函数：根据状态码返回颜色
func statusCodeColor(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "#27ae60"
	case code >= 400 && code < 500:
		return "#e74c3c"
	case code >= 500:
		return "#f39c12"
	default:
		return "#95a5a6"
	}
}
