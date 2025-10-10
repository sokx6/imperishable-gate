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
            font-family: 'Segoe UI', 'Microsoft YaHei', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(to bottom, #fff5f8 0%%, #fef8fa 100%%);
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 650px;
            margin: 40px auto;
            background: white;
            border-radius: 12px;
            overflow: hidden;
            box-shadow: 0 8px 24px rgba(255, 182, 193, 0.2);
            border: 2px solid #ffd4e5;
        }
        .header {
            background: linear-gradient(to right, #ffb8d1 0%%, #ffd4e5 50%%, #ffb8d1 100%%);
            color: #4a3842;
            padding: 35px 25px;
            text-align: center;
            border-bottom: 3px double #ffb8d1;
        }
        .header h1 {
            margin: 0;
            font-size: 32px;
            font-weight: 700;
            letter-spacing: 2px;
            text-shadow: 2px 2px 4px rgba(255, 255, 255, 0.3);
        }
        .header .subtitle {
            margin: 8px 0 0 0;
            font-size: 15px;
            opacity: 0.85;
            font-weight: 500;
        }
        .header .stage {
            margin: 5px 0 0 0;
            font-size: 12px;
            opacity: 0.7;
            font-style: italic;
        }
        .content {
            padding: 40px 35px;
            background: linear-gradient(to bottom, #ffffff 0%%, #fafcff 100%%);
        }
        .greeting {
            color: #5a4a52;
            font-size: 16px;
            margin-bottom: 20px;
            line-height: 1.8;
            border-left: 4px solid #ffb8d1;
            padding-left: 15px;
            background: #fffafc;
            padding: 15px 15px 15px 20px;
            border-radius: 4px;
        }
        .code-box {
            background: linear-gradient(to right, #ff9eba 0%%, #ffb8d1 50%%, #ff9eba 100%%);
            color: white;
            font-size: 42px;
            font-weight: bold;
            letter-spacing: 10px;
            text-align: center;
            padding: 30px;
            margin: 30px 0;
            border-radius: 10px;
            box-shadow: 0 6px 20px rgba(255, 184, 209, 0.35);
            border: 2px solid #ffcce0;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.15);
        }
        .info-box {
            background: #fffafc;
            border-left: 4px solid #ffb8d1;
            padding: 18px 22px;
            margin: 25px 0;
            border-radius: 6px;
            box-shadow: 0 2px 8px rgba(255, 184, 209, 0.1);
        }
        .info-box p {
            margin: 8px 0;
            color: #5a4a52;
            font-size: 14px;
        }
        .info-box strong {
            color: #d4748e;
            font-weight: 600;
        }
        .warning {
            background: #fff8f8;
            border-left: 4px solid #d47474;
            padding: 15px 20px;
            margin: 25px 0;
            border-radius: 6px;
            color: #8a5a5a;
            font-size: 13px;
            line-height: 1.6;
        }
        .footer {
            background: linear-gradient(to bottom, #fffafc 0%%, #fff5f8 100%%);
            padding: 25px;
            text-align: center;
            color: #a08a95;
            font-size: 12px;
            border-top: 2px solid #ffd4e5;
        }
        .footer .quote {
            font-style: italic;
            color: #d4748e;
            margin-bottom: 10px;
            font-size: 13px;
        }
        .footer .signature {
            margin-top: 15px;
            color: #c0a8b0;
            font-size: 11px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Imperishable Gate</h1>
            <div class="subtitle">不朽之门 · 白玉楼链接管理系统</div>
            <div class="stage">Stage 6 | 冥界大小姐の亡骸</div>
        </div>
        <div class="content">
            <div class="greeting">
                您好，访客：<br/>
                您正在进行邮箱验证操作。作为白玉楼的庭师，我已为您准备好通行验证码。<br/>
                请在 15 分钟内使用以下验证码完成身份确认。
            </div>
            
            <div class="code-box">%s</div>
            
            <div class="info-box">
                <p><strong>验证码有效期：</strong>15 分钟</p>
                <p><strong>使用方式：</strong>请在客户端或网页中输入此验证码</p>
                <p><strong>验证场景：</strong>注册账户 / 登录验证 / 修改邮箱等安全操作</p>
            </div>
            
            <div class="warning">
                <strong>安全提示：</strong><br/>
                若非本人操作，请忽略此邮件。验证码关乎账户安全，请妥善保管，切勿告知他人。<br/>
                白玉楼的门扉只为真正的主人开启。
            </div>
        </div>
        <div class="footer">
            <div class="quote">"想通过这扇门？先证明你的身份！"</div>
            <p>Imperishable Gate - 基于东方 Project 的链接管理系统</p>
            <div class="signature">
                白玉楼庭师 敬上<br/>
                本邮件由系统自动发送，请勿回复
            </div>
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
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: 'Segoe UI', 'Microsoft YaHei', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(to bottom, #fff5f8 0%%, #fef8fa 100%%);
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 700px;
            margin: 40px auto;
            background: white;
            border-radius: 12px;
            overflow: hidden;
            box-shadow: 0 8px 24px rgba(255, 182, 193, 0.2);
            border: 2px solid #ffd4e5;
        }
        .header {
            background: linear-gradient(to right, #ffb8d1 0%%, #ffd4e5 50%%, #ffb8d1 100%%);
            color: #4a3842;
            padding: 35px 25px;
            text-align: center;
            border-bottom: 3px double #ffb8d1;
        }
        .header h1 {
            margin: 0;
            font-size: 28px;
            font-weight: 700;
            letter-spacing: 2px;
        }
        .header .subtitle {
            margin: 8px 0 0 0;
            font-size: 14px;
            opacity: 0.85;
        }
        .header .stage {
            margin: 5px 0 0 0;
            font-size: 12px;
            opacity: 0.7;
            font-style: italic;
        }
        .content {
            padding: 35px 30px;
            background: linear-gradient(to bottom, #ffffff 0%%, #fafcff 100%%);
        }
        .alert-box {
            background: linear-gradient(to right, #fff0f5 0%%, #fffafc 100%%);
            border-left: 5px solid #ffb8d1;
            padding: 20px;
            margin-bottom: 25px;
            border-radius: 8px;
            box-shadow: 0 2px 8px rgba(255, 184, 209, 0.1);
        }
        .alert-box h3 {
            margin: 0 0 10px 0;
            color: #d4748e;
            font-size: 20px;
        }
        .alert-box .url {
            margin: 10px 0;
            word-break: break-all;
        }
        .alert-box .url strong {
            color: #5a4a52;
        }
        .alert-box .url a {
            color: #d4748e;
            text-decoration: none;
            font-weight: 500;
        }
        .alert-box .url a:hover {
            text-decoration: underline;
        }
        .comparison-table {
            width: 100%%;
            border-collapse: collapse;
            margin: 20px 0;
            background: #fffafc;
            border-radius: 8px;
            overflow: hidden;
            box-shadow: 0 2px 8px rgba(255, 184, 209, 0.08);
        }
        .comparison-table thead {
            background: linear-gradient(to right, #ffb8d1 0%%, #ffd4e5 100%%);
            color: #4a3842;
        }
        .comparison-table th {
            padding: 15px 12px;
            text-align: left;
            font-weight: 600;
            font-size: 14px;
        }
        .comparison-table td {
            padding: 12px;
            border-bottom: 1px solid #ffe8f0;
            color: #5a4a52;
            font-size: 14px;
        }
        .comparison-table tr:last-child td {
            border-bottom: none;
        }
        .comparison-table .label {
            font-weight: 600;
            color: #d4748e;
            width: 80px;
        }
        .comparison-table .old-value {
            color: #999;
            font-style: italic;
        }
        .comparison-table .new-value {
            color: #6ba56b;
            font-weight: 500;
        }
        .footer {
            background: linear-gradient(to bottom, #fffafc 0%%, #fff5f8 100%%);
            padding: 25px;
            text-align: center;
            color: #a08a95;
            font-size: 12px;
            border-top: 2px solid #ffd4e5;
        }
        .footer .quote {
            font-style: italic;
            color: #d4748e;
            margin-bottom: 10px;
            font-size: 13px;
        }
        .footer .signature {
            margin-top: 15px;
            color: #c0a8b0;
            font-size: 11px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Imperishable Gate</h1>
            <div class="subtitle">不朽之门 · 白玉楼链接管理系统</div>
            <div class="stage">Stage 4 | 雪上の櫻花結界 - 智能监控</div>
        </div>
        <div class="content">
            <div class="alert-box">
                <h3>网站内容已更新</h3>
                <p style="color: #5a4a52; line-height: 1.8; margin: 10px 0;">
                    您好，白玉楼的访客：<br/>
                    作为庭师，我发现您关注的网站发生了变化。以下是详细的变更记录，请查收。
                </p>
                <div class="url">
                    <strong>监控网址：</strong><a href="%s" target="_blank">%s</a>
                </div>
            </div>

            <table class="comparison-table">
                <thead>
                    <tr>
                        <th>项目</th>
                        <th>旧内容</th>
                        <th>新内容</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td class="label">标题</td>
                        <td class="old-value">%s</td>
                        <td class="new-value">%s</td>
                    </tr>
                    <tr>
                        <td class="label">描述</td>
                        <td class="old-value" style="white-space: pre-wrap;">%s</td>
                        <td class="new-value" style="white-space: pre-wrap;">%s</td>
                    </tr>
                    <tr>
                        <td class="label">关键词</td>
                        <td class="old-value">%s</td>
                        <td class="new-value">%s</td>
                    </tr>
                    <tr>
                        <td class="label">状态码</td>
                        <td style="color: %s; font-weight: bold;">%d</td>
                        <td style="color: %s; font-weight: bold;">%d</td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="footer">
            <div class="quote">"时刻关注你在意的网站变化"</div>
            <p>Imperishable Gate - 基于东方 Project 的链接管理系统</p>
            <div class="signature">
                白玉楼庭师 敬上<br/>
                本邮件由监控系统自动发送，请勿回复
            </div>
        </div>
    </div>
</body>
</html>
	`,
		changedUrl, changedUrl,
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
