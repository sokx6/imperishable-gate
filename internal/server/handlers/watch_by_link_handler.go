package handlers

import (
	"errors"
	"net/http"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"github.com/labstack/echo/v4"
)

func WatchByUrlHandler(c echo.Context) error {
	var req request.WatchByUrlRequest
	if err := c.Bind(&req); err != nil || req.Url == "" {
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := service.Watch(req.Url, userId, req.Watch); err != nil {
		if errors.Is(err, service.ErrLinkNotFound) {
			return response.LinkNotFoundResponse
		}
		return response.DatabaseErrorResponse
	}
	if req.Watch {
		return c.JSON(http.StatusOK, response.WatchSuccessResponse)
	}
	return c.JSON(http.StatusOK, response.UnwatchSuccessResponse)
}
