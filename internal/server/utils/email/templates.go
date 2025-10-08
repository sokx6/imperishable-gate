package email

import "fmt"

// GetVerificationEmailTemplate 获取验证邮件的 HTML 模板
func GetVerificationEmailTemplate(code string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 600px;
            margin: 40px auto;
            background: white;
            border-radius: 10px;
            overflow: hidden;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            color: white;
            padding: 30px 20px;
            text-align: center;
        }
        .header h1 {
            margin: 0;
            font-size: 28px;
            font-weight: 600;
        }
        .header p {
            margin: 10px 0 0 0;
            font-size: 14px;
            opacity: 0.9;
        }
        .content {
            padding: 40px 30px;
        }
        .content h2 {
            color: #333;
            font-size: 20px;
            margin-top: 0;
        }
        .content p {
            color: #666;
            line-height: 1.6;
            font-size: 15px;
        }
        .code-box {
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            color: white;
            font-size: 36px;
            font-weight: bold;
            letter-spacing: 8px;
            text-align: center;
            padding: 25px;
            margin: 30px 0;
            border-radius: 8px;
            box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
        }
        .info-box {
            background: #f8f9fa;
            border-left: 4px solid #667eea;
            padding: 15px 20px;
            margin: 20px 0;
            border-radius: 4px;
        }
        .info-box p {
            margin: 5px 0;
            color: #555;
        }
        .info-box strong {
            color: #667eea;
        }
        .footer {
            background: #f8f9fa;
            padding: 20px;
            text-align: center;
            color: #999;
            font-size: 12px;
            border-top: 1px solid #e0e0e0;
        }
        .warning {
            color: #e74c3c;
            font-size: 13px;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🔐 Imperishable Gate</h1>
            <p>链接管理系统</p>
        </div>
        <div class="content">
            <h2>欢迎注册！</h2>
            <p>感谢您注册 Imperishable Gate 链接管理系统。</p>
            <p>请使用以下验证码完成邮箱验证：</p>
            
            <div class="code-box">%s</div>
            
            <div class="info-box">
                <p><strong>⏰ 有效期：</strong>15分钟</p>
                <p><strong>💡 提示：</strong>如果您在客户端中，请直接输入此验证码</p>
            </div>
            
            <p class="warning">⚠️ 如果这不是您的操作，请忽略此邮件。为了您的账户安全，请勿将验证码告诉他人。</p>
        </div>
        <div class="footer">
            <p>© 2025 Imperishable Gate. All rights reserved.</p>
            <p>这是一封自动发送的邮件，请勿回复。</p>
        </div>
    </div>
</body>
</html>
    `, code)
}

// GetWebsiteChangeEmailTemplate 获取网站变更通知邮件的 HTML 模板
func GetWebsiteChangeEmailTemplate(
	oldTitle, oldDesc, oldKeywords,
	newTitle, newDesc, newKeyword,
	changedUrl string,
	oldStatusCode, newStatusCode int) string {

	return fmt.Sprintf(`
<html>
<body style="font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; color: #333; line-height: 1.6; max-width: 700px; margin: auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 12px; background-color: #f9f9fb;">
	<div style="text-align: center; margin-bottom: 20px;">
		<h2 style="color: #4a90e2;">🚀 网站监控小助手</h2>
		<p style="color: #666; font-size: 14px;">第一时间为您推送重要变化 💬</p>
		<hr style="border: 1px dashed #ccc;" />
	</div>

	<h3 style="color: #2c3e50;">📌 页面已更新</h3>
	<p><strong>📍 网址:</strong> <a href="%s" style="color: #4a90e2; text-decoration: none;">访问此页</a></p>

	<table width="100%%" style="background-color: #f1f8ff; padding: 15px; border-radius: 8px; border: 1px solid #d0ebff; margin: 10px 0;">
		<tr>
			<td style="padding: 8px;"><strong>📌 项目</strong></td>
			<td style="padding: 8px;"><strong>旧内容</strong></td>
			<td style="padding: 8px;"><strong>新内容 🎉</strong></td>
		</tr>
		<tr>
			<td style="padding: 8px;">📄 标题</td>
			<td style="padding: 8px; color: #999;"><i>%s</i></td>
			<td style="padding: 8px; color: #27ae60; font-weight: bold;">%s</td>
		</tr>
		<tr>
			<td style="padding: 8px;">📝 描述</td>
			<td style="padding: 8px; color: #999; white-space: pre-wrap;"><i>%s</i></td>
			<td style="padding: 8px; color: #27ae60; white-space: pre-wrap;">%s</td>
		</tr>
		<tr>
			<td style="padding: 8px;">🔖 关键词</td>
			<td style="padding: 8px; color: #999;"><i>%s</i></td>
			<td style="padding: 8px; color: #27ae60;">%s</td>
		</tr>
		<tr>
			<td style="padding: 8px;">🔧 状态码</td>
			<td style="padding: 8px; color: %s;"><b>%d</b></td>
			<td style="padding: 8px; color: %s;"><b>%d</b></td>
		</tr>
	</table>

	<div style="margin-top: 20px; text-align: center; font-size: 14px; color: #777;">
		<p>🔔 这是自动通知，请勿回复本邮件。<br/>
		如有疑问，欢迎联系支持团队 ❤️</p>
		<small>Powered by Your Site Monitor ✨</small>
	</div>
</body>
</html>
	`,
		changedUrl,
		HtmlEscape(oldTitle),
		HtmlEscape(newTitle),
		HtmlEscape(TruncateText(oldDesc, 100)),
		HtmlEscape(TruncateText(newDesc, 100)),
		HtmlEscape(TruncateText(oldKeywords, 80)),
		HtmlEscape(TruncateText(newKeyword, 80)),
		StatusCodeColor(oldStatusCode),
		oldStatusCode,
		StatusCodeColor(newStatusCode),
		newStatusCode,
	)
}
