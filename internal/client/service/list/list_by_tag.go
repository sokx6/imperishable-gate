package list

import (
	"net/http"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

func ListByTag(addr string, accessToken string, tag string, page int, pageSize int) (response.ListResponse, error) {
	client := utils.NewAPIClient(addr, accessToken)

	var respBody response.ListResponse
	if err := client.DoRequest(http.MethodGet, "/api/v1/tags/"+tag, nil, &respBody); err != nil {
		return response.ListResponse{}, err
	}

	return respBody, nil
}
