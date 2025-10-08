package watch

import (
	"fmt"

	"imperishable-gate/internal/client/utils"
)

// HandleWatchByLink 通过链接 URL 设置监控
func HandleWatchByLink(link string, watch bool, addr, accessToken string) error {

	// 调用 watch by link 服务
	return WatchByLink(link, watch, addr, accessToken)
}

// HandleWatchByName 通过名称设置监控
func HandleWatchByName(name string, watch bool, addr, accessToken string) error {

	// 调用 watch by name 服务
	return WatchByName(name, watch, addr, accessToken)
}

// 打印成功消息
func printSuccessMessage(watch bool, identifier string) {
	action := "unwatching"
	if watch {
		action = "watching"
	}
	fmt.Printf("Successfully set %s for: %s\n", action, identifier)
}

// 处理通用错误
func handleWatchError(err error, watch bool) error {
	action := "unwatch"
	if watch {
		action = "watch"
	}
	return fmt.Errorf("failed to %s link: %w", action, err)
}

// CreateAPIClient 创建 API 客户端
func CreateAPIClient(addr, token string) *utils.APIClient {
	return utils.NewAPIClient(addr, token)
}
