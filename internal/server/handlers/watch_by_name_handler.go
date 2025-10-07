package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func WatchByNameHandler(c echo.Context) error {
	var req request.WatchByNameRequest
	if err := c.Bind(&req); err != nil || req.Name == "" {
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	linkUrl := utils.GetLinkUrlByName(req.Name, userId)
	if err := service.Watch(linkUrl, userId, req.Watch); err != nil {
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
