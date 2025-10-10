package links

import (
	"errors"
	"net/http"

	"imperishable-gate/internal/server/service/common"
	linksService "imperishable-gate/internal/server/service/links"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"github.com/labstack/echo/v4"
)

func WatchByUrlHandler(c echo.Context) error {
	var req request.WatchByUrlRequest
	if err := c.Bind(&req); err != nil || req.Url == "" {
		logger.Warning("Invalid watch/unwatch request: empty URL or invalid format")
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := linksService.Watch(req.Url, userId, req.Watch); err != nil {
		if errors.Is(err, common.ErrLinkNotFound) {
			logger.Warning("Watch/Unwatch failed: link %s not found for user %d", req.Url, userId)
			return response.LinkNotFoundResponse
		}
		logger.Error("Database error while watching/unwatching link %s for user %d: %v", req.Url, userId, err)
		return response.DatabaseErrorResponse
	}
	if req.Watch {
		logger.Success("Link %s watched successfully for user %d", req.Url, userId)
		return c.JSON(http.StatusOK, response.WatchSuccessResponse)
	}
	logger.Success("Link %s unwatched successfully for user %d", req.Url, userId)
	return c.JSON(http.StatusOK, response.UnwatchSuccessResponse)
}
