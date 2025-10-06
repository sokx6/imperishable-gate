package add

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

func AddTagsByLink(link string, tags []string, addr string, accessToken string) error {
	client := utils.NewAPIClient(addr, accessToken)

	reqBody := request.AddRequest{
		Link: link,
		Tags: tags,
	}

	var respBody response.Response
	if err := client.DoRequest(http.MethodPost, "/api/v1/tags", reqBody, &respBody); err != nil {
		return err
	}

	return nil

}
