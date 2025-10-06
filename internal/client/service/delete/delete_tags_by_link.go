package delete

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func DeleteTagsByLink(url string, userId int, tags []string, addr string, accessToken string) error {
	client := utils.NewAPIClient(addr, accessToken)
	reqBody := request.DeleteRequest{
		Url:  url,
		Tags: tags,
	}
	var resp response.Response
	if err := client.DoRequest("PATCH", "/api/v1/links/by-url/tags/remove", reqBody, &resp); err != nil {
		return err
	}
	return nil
}
