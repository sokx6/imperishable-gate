package delete

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

func DeleteNamesByLink(link string, names []string, addr string, accessToken string) error {
	// 规范化 URL（如果缺少协议则自动添加 https://）
	link = utils.NormalizeURL(link)

	client := utils.NewAPIClient(addr, accessToken)
	reqBody := request.DeleteRequest{
		Url:   link,
		Names: names,
	}

	var resp response.Response
	if err := client.DoRequest(http.MethodPatch, "/api/v1/links/names/remove", reqBody, &resp); err != nil {
		return err
	}
	return nil
}
