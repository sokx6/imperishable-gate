package add

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

func AddNames(link string, names []string, addr string, accessToken string) error {
	// 创建 API 客户端
	client := utils.NewAPIClient(addr, accessToken)

	// 构建请求体
	reqBody := request.AddRequest{
		Link:  link,
		Names: names,
	}

	// 发送请求并处理响应
	var respBody response.Response
	if err := client.DoRequest(http.MethodPost, "/api/v1/names", reqBody, &respBody); err != nil {
		// 处理请求错误
		return err
	}

	return nil
}
