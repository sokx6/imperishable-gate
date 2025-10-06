package delete

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func DeleteNamesByLink(link string, names []string, addr string, accessToken string) error {
	client := utils.NewAPIClient(addr, accessToken)
	reqBody := request.DeleteRequest{
		Names: names,
	}

	var resp response.Response
	if err := client.DoRequest("DELETE", "/api/v1/names?link="+link, reqBody, &resp); err != nil {
		return err
	}
	return nil
}
