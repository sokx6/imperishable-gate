package list

import (
	"net/http"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

func ListByName(addr string, accessToken string, name string, page int, pageSize int) (response.ListByNameResponse, error) {
	// 创建 API 客户端
	client := utils.NewAPIClient(addr, accessToken)

	// 发送请求并处理响应
	var respBody response.ListByNameResponse
	if err := client.DoRequest(http.MethodGet, "/api/v1/names/"+name, nil, &respBody); err != nil {
		return response.ListByNameResponse{}, err
	}

	return respBody, nil
}
