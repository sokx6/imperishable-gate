package delete

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

func DeleteNamesByLink(link string, names []string, addr string, accessToken string) error {
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
