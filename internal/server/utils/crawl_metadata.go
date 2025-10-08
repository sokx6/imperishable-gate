package utils

import (
	"imperishable-gate/internal/server/utils/logger"
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

func CrawlMetadata(rawURL string) (title, desc, keywords string, statusCode int, err error) {
	title, desc, keywords = "", "", ""
	statusCode = 0
	err = nil

	logger.Debug("Crawling metadata for URL: %s", rawURL)

	// 创建带 cookiejar 的 HTTP 客户端
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		logger.Error("Failed to create cookie jar for URL %s: %v", rawURL, err)
		return
	}
	client := &http.Client{Jar: jar}
	// 发送带模拟浏览器头部的请求
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		logger.Error("Failed to create request for URL %s: %v", rawURL, err)
		return
	}
	// 模拟浏览器头部
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml;q=0.9")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		if err != nil {
			logger.Warning("Failed to fetch URL %s: %v", rawURL, err)
		} else {
			logger.Warning("Non-OK status code %d for URL: %s", resp.StatusCode, rawURL)
		}
		return
	}
	defer resp.Body.Close()
	// 获取状态码
	statusCode = resp.StatusCode
	// 用 goquery 解析
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error("Failed to parse HTML for URL %s: %v", rawURL, err)
		return
	}
	title = doc.Find("title").Text()

	desc, _ = doc.Find(`meta[name="description"]`).Attr("content")
	keywords, _ = doc.Find(`meta[name="keywords"]`).Attr("content")

	logger.Debug("Metadata crawled successfully for URL: %s (title: %s)", rawURL, title)
	return
}
