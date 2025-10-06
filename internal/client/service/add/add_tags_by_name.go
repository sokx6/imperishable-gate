package add

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
	"path"
)

func AddTagsByName(name string, tags []string, addr string, accessToken string) error {
	client := utils.NewAPIClient(addr, accessToken)

	reqBody := request.AddRequest{
		Name: name,
		Tags: tags,
	}

	var respBody response.Response
	apiPath := path.Join("/api/v1/name", name, "tags")
	if err := client.DoRequest(http.MethodPost, apiPath, reqBody, &respBody); err != nil {
		return err
	}

	return nil
}
