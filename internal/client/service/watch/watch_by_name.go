package watch

import (
	"fmt"
	"net/http"

	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

// WatchByName 通过名称设置监控状态
func WatchByName(name string, watch bool, addr, token string) error {
	client := CreateAPIClient(addr, token)

	// 构建请求路径（名称作为路径参数）
	path := "/api/v1/name/watch"

	// 构建请求体
	reqBody := request.WatchByNameRequest{
		Name:  name,
		Watch: watch,
	}

	// 调用 API
	var resp response.Response
	err := client.DoRequest(http.MethodPatch, path, reqBody, &resp)
	if err != nil {
		return handleWatchError(err, watch)
	}

	// 打印成功消息
	printSuccessMessage(watch, name)
	fmt.Printf("Message: %s\n", resp.Message)

	return nil
}
