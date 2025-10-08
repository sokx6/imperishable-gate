package utils

import "strings"

// NormalizeURL 规范化 URL，如果缺少协议则自动添加 https://
func NormalizeURL(url string) string {
	if url == "" {
		return url
	}

	// 如果 URL 没有协议前缀，自动添加 https://
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "https://" + url
	}

	return url
}

// NormalizeURLs 批量规范化 URL 列表
func NormalizeURLs(urls []string) []string {
	normalized := make([]string, len(urls))
	for i, url := range urls {
		normalized[i] = NormalizeURL(url)
	}
	return normalized
}
