package tags

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service/common"
	tagsService "imperishable-gate/internal/server/service/tags"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func DeleteTagsByLinkHandler(c echo.Context) error {
	var req request.DeleteRequest
	if err := c.Bind(&req); err != nil || req.Url == "" || req.Tags == nil || len(req.Tags) == 0 {
		logger.Warning("Invalid delete tags request: empty URL or tags")
		return response.InvalidRequestResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	if err := tagsService.DeleteTagsByLink(req.Url, userId, req.Tags); err != nil {
		if errors.Is(err, common.ErrLinkNotFound) {
			logger.Warning("Link not found: %s", req.Url)
			return response.LinkNotFoundResponse
		}
		if errors.Is(err, common.ErrInvalidRequest) {
			logger.Warning("Invalid delete tags request: %v", err)
			return response.InvalidRequestResponse
		}
		logger.Error("Database error while deleting tags: %v", err)
		return response.DatabaseErrorResponse
	}

	logger.Success("Tags deleted successfully for link %s by user %d", req.Url, userId)
	return c.JSON(http.StatusOK, response.DeleteTagsByLinkSuccessResponse)
}
