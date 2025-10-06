package handlers

import (
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func WatchByLinkHandler(c echo.Context) error {
	var req request.WatchRequest
	if err := c.Bind(&req); err != nil || req.Url == "" {
		return response.InvalidRequestResponse
	}
	userId := c.Get("userId").(uint)
	if err := service.Watch(req.Url, userId, req.Watch); err != nil {
		if err == service.ErrLinkNotFound {
			return response.LinkNotFoundResponse
		}
		return response.DatabaseErrorResponse
	}
	if req.Watch {
		return c.JSON(http.StatusOK, response.WatchSuccessResponse)
	}
	return c.JSON(http.StatusOK, response.UnwatchSuccessResponse)
}
