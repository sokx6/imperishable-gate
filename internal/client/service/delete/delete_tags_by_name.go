package delete

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

func DeleteTagsByName(name string, tags []string, addr string, accessToken string) error {
	client := utils.NewAPIClient(addr, accessToken)

	reqBody := request.DeleteRequest{
		Tags: tags,
	}
	path := "/api/v1/" + name + "/tags/remove"
	var resp response.Response
	if err := client.DoRequest(http.MethodPatch, path, reqBody, &resp); err != nil {
		return err
	}
	return nil
}
