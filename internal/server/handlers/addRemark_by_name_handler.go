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

func AddRemarkByNameHandler(c echo.Context) error {
	var req request.AddRequest
	name := c.Param("name")
	// 检查请求的有效性
	if err := c.Bind(&req); err != nil || name == "" || req.Remark == "" {
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := service.AddRemarkByName(name, userId, req.Remark); err != nil {
		if errors.Is(err, service.ErrNameNotFound) {
			return response.NameNotFoundResponse
		}
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.AddRemarkByNameSuccessResponse)
}
