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

func AddTagsByNameHandler(c echo.Context) error {
	name := c.Param("name")
	var req request.AddRequest
	if err := c.Bind(&req); err != nil || req.Tags == nil || len(req.Tags) == 0 {
		return response.InvalidRequestResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := service.AddTagsByName(name, userId, req.Tags); err != nil {
		if errors.Is(err, service.ErrNameNotFound) {
			return response.NameNotFoundResponse
		}
		return response.DatabaseErrorResponse
	}
	return c.JSON(http.StatusOK, response.AddTagsByNameSuccessResponse)

}
