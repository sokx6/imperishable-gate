package list

import (
	"fmt"
	"net/http"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

func ListLinks(addr string, accessToken string, page int, pageSize int) (response.Response, error) {
	// 构建请求路径和查询参数
	pathWithQuery := fmt.Sprintf("/api/v1/links?page=%d&page_size=%d", page, pageSize)

	// 创建 API 客户端
	client := utils.NewAPIClient(addr, accessToken)

	// 发送请求并处理响应
	var respBody response.Response
	if err := client.DoRequest(http.MethodGet, pathWithQuery, nil, &respBody); err != nil {
		return response.Response{}, err
	}

	return respBody, nil
}
