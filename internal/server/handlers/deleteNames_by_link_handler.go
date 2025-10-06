package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func DeleteNamesByLinkHandler(c echo.Context) error {
	var req request.DeleteRequest
	url := c.QueryParam("url")
	if err := c.Bind(&req); err != nil || url == "" || req.Names == nil || len(req.Names) == 0 {
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := service.DeleteNamesByLink(url, userId, req.Names); err != nil {
		if err == service.ErrLinkNotFound {
			return response.LinkNotFoundResponse
		}
		if err == service.ErrInvalidRequest {
			return response.InvalidRequestResponse
		}
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.DeleteNamesByLinkSuccessResponse)
}
