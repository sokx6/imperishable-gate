package add

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
	"path"
)

func AddRemarkByName(name string, remark string, addr string, accessToken string) error {
	// 创建 API 客户端
	client := utils.NewAPIClient(addr, accessToken)

	// 构建请求体
	reqBody := request.AddRequest{
		Name:   name,
		Remark: remark,
	}

	// 发送请求并处理响应
	var respBody response.Response
	path := path.Join("/api/v1/name", name, "remark")
	if err := client.DoRequest(http.MethodPost, path, reqBody, &respBody); err != nil {
		// 处理请求错误
		return err
	}

	return nil
}
