package list

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

func ListByTag(addr string, accessToken string, tag string, page int, pageSize int) (response.ListResponse, error) {
	// 解析基础地址
	baseURL, err := url.Parse(addr)
	if err != nil {
		return response.ListResponse{}, fmt.Errorf("invalid address: %w", err)
	}

	// 构建请求路径：/api/v1/tags/{tag}
	baseURL.Path = path.Join(baseURL.Path, "/api/v1/tags", tag)

	// 添加查询参数
	q := baseURL.Query()
	q.Set("page", fmt.Sprintf("%d", page))
	q.Set("pageSize", fmt.Sprintf("%d", pageSize))
	baseURL.RawQuery = q.Encode()
	// 创建 API 客户端
	client := utils.NewAPIClient(addr, accessToken)

	// 发送请求并处理响应
	var respBody response.ListResponse
	if err := client.DoRequest(http.MethodGet, baseURL.String(), nil, &respBody); err != nil {
		return response.ListResponse{}, err
	}

	return respBody, nil
}
