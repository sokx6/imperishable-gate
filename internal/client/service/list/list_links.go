package list

import (
	"net/http"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

func ListLinks(addr string, accessToken string, page int, pageSize int) (response.ListResponse, error) {
	// 创建 API 客户端
	client := utils.NewAPIClient(addr, accessToken)

	// 发送请求并处理响应
	var respBody response.ListResponse
	if err := client.DoRequest(http.MethodGet, "/api/v1/links", nil, &respBody); err != nil {
		return response.ListResponse{}, err
	}

	return respBody, nil
}
