package list

import (
	"fmt"
	"net/http"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

func ListByName(addr string, accessToken string, name string) (response.Response, error) {
	// 构建请求路径
	path := fmt.Sprintf("/api/v1/names/%s", name)

	// 创建 API 客户端
	client := utils.NewAPIClient(addr, accessToken)

	// 发送请求并处理响应
	var respBody response.Response
	if err := client.DoRequest(http.MethodGet, path, nil, &respBody); err != nil {
		return response.Response{}, err
	}

	return respBody, nil
}
