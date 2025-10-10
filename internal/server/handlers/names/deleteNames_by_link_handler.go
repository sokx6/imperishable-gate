package names

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service/common"
	namesService "imperishable-gate/internal/server/service/names"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func DeleteNamesByLinkHandler(c echo.Context) error {
	var req request.DeleteRequest
	if err := c.Bind(&req); err != nil || req.Url == "" || req.Names == nil || len(req.Names) == 0 {
		logger.Warning("Invalid delete names request: empty link or names")
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := namesService.DeleteNamesByLink(req.Url, userId, req.Names); err != nil {
		if errors.Is(err, common.ErrLinkNotFound) {
			logger.Warning("Delete names failed: link %s not found for user %d", req.Url, userId)
			return response.LinkNotFoundResponse
		}
		if errors.Is(err, common.ErrInvalidRequest) {
			logger.Warning("Invalid delete names request: %v", req.Names)
			return response.InvalidRequestResponse
		}
		logger.Error("Database error while deleting names for link %s and user %d: %v", req.Url, userId, err)
		return response.DatabaseErrorResponse
	}

	logger.Success("Names %v deleted successfully for link %s by user %d", req.Names, req.Url, userId)
	return c.JSON(http.StatusOK, response.DeleteNamesByLinkSuccessResponse)
}
