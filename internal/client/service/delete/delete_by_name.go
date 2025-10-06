package delete

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
	"net/http"
	"path"
)

func DeleteByName(name string, addr string, accessToken string) error {
	client := utils.NewAPIClient(addr, accessToken)

	respBody := response.Response{}
	path := path.Join("/api/v1/name", name)
	if err := client.DoRequest(http.MethodDelete, path, nil, &respBody); err != nil {
		return err
	}

	return nil
}
