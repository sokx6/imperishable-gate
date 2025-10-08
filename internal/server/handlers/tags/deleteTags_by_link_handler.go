package tags

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	tagsService "imperishable-gate/internal/server/service/tags"
	"imperishable-gate/internal/server/service/common"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func DeleteTagsByLinkHandler(c echo.Context) error {
	var req request.DeleteRequest
	if err := c.Bind(&req); err != nil || req.Url == "" || req.Tags == nil || len(req.Tags) == 0 {
		return response.InvalidRequestResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	if err := tagsService.DeleteTagsByLink(req.Url, userId, req.Tags); err != nil {
		if errors.Is(err, common.ErrLinkNotFound) {
			return response.LinkNotFoundResponse
		}
		if errors.Is(err, common.ErrInvalidRequest) {
			return response.InvalidRequestResponse
		}
		return response.DatabaseErrorResponse
	}
	return c.JSON(http.StatusOK, response.DeleteTagsByLinkSuccessResponse)
}
