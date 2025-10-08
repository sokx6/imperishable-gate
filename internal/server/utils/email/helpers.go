package email

import "html/template"

// HtmlEscape 转义 HTML 特殊字符，防止 XSS 注入或格式错乱
func HtmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

// TruncateText 截断过长的文本，避免表格太宽
func TruncateText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}
	return text[:maxLength] + "..."
}

// StatusCodeColor 根据状态码返回颜色
func StatusCodeColor(code int) string {
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
