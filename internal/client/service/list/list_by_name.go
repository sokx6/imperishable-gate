package list

import (
	"net/http"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

func ListByName(addr string, accessToken string, name string, page int, pageSize int) (response.ListByNameResponse, error) {
	client := utils.NewAPIClient(addr, accessToken)

	var respBody response.ListByNameResponse
	if err := client.DoRequest(http.MethodGet, "/api/v1/names/"+name, nil, &respBody); err != nil {
		return response.ListByNameResponse{}, err
	}

	return respBody, nil
}
