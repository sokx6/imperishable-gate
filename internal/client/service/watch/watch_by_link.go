package watch

import (
	"fmt"
	"net/http"

	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

// WatchByLink 通过链接 URL 设置监控状态
func WatchByLink(link string, watch bool, addr, token string) error {
	client := CreateAPIClient(addr, token)

	// 构建请求体
	reqBody := request.WatchByUrlRequest{
		Url:   link,
		Watch: watch,
	}

	// 调用 API
	var resp response.Response
	err := client.DoRequest(http.MethodPatch, "/api/v1/links/watch", reqBody, &resp)
	if err != nil {
		return handleWatchError(err, watch)
	}

	// 打印成功消息
	printSuccessMessage(watch, link)
	fmt.Printf("Message: %s\n", resp.Message)

	return nil
}
