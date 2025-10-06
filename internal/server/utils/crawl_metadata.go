package utils

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

func CrawlMetadata(rawURL string) (title, desc, keywords string, statusCode int) {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client := &http.Client{Jar: jar}
	// 发送带模拟浏览器头部的请求
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", "", "", 0
	}
	// 模拟浏览器头部
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml;q=0.9")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", "", "", 0
	}
	defer resp.Body.Close()
	// 获取状态码
	statusCode = resp.StatusCode
	// 用 goquery 解析
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	title = doc.Find("title").Text()

	desc, _ = doc.Find(`meta[name="description"]`).Attr("content")
	keywords, _ = doc.Find(`meta[name="keywords"]`).Attr("content")

	return
}
