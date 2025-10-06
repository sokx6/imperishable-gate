package add

import (
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
	"path"
)

func AddRemarkByName(name string, remark string, addr string, accessToken string) error {
	client := utils.NewAPIClient(addr, accessToken)

	reqBody := request.AddRequest{
		Name:   name,
		Remark: remark,
	}

	var respBody response.Response
	path := path.Join("/api/v1/name", name, "remarks")
	if err := client.DoRequest(http.MethodPost, path, reqBody, &respBody); err != nil {
		return err
	}

	return nil
}
