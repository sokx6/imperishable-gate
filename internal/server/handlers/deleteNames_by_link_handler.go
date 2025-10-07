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

func DeleteNamesByLinkHandler(c echo.Context) error {
	var req request.DeleteRequest
	if err := c.Bind(&req); err != nil || req.Url == "" || req.Names == nil || len(req.Names) == 0 {
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := service.DeleteNamesByLink(req.Url, userId, req.Names); err != nil {
		if errors.Is(err, service.ErrLinkNotFound) {
			return response.LinkNotFoundResponse
		}
		if errors.Is(err, service.ErrInvalidRequest) {
			return response.InvalidRequestResponse
		}
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.DeleteNamesByLinkSuccessResponse)
}
