package list

import (
	"net/http"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

func ListLinks(addr string, accessToken string, page int, pageSize int) (response.ListResponse, error) {
	client := utils.NewAPIClient(addr, accessToken)

	var respBody response.ListResponse
	if err := client.DoRequest(http.MethodGet, "/api/v1/links", nil, &respBody); err != nil {
		return response.ListResponse{}, err
	}

	return respBody, nil
}
