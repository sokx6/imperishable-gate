package add

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

func AddNames(link string, names []string, addr string, accessToken string) error {
	client := utils.NewAPIClient(addr, accessToken)

	reqBody := request.AddRequest{
		Link:  link,
		Names: names,
	}

	var respBody response.Response
	if err := client.DoRequest(http.MethodPost, "/api/v1/names", reqBody, &respBody); err != nil {
		return err
	}

	return nil
}
