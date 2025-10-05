package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
)

func AddRemarkByNameHandler(c echo.Context) error {
	var req types.AddRemarkByNameRequest
	name := c.Param("name")
	// 检查请求的有效性
	if err := c.Bind(&req); err != nil || name == "" || req.Remark == "" {
		return c.JSON(400, types.InvalidRequestResponse)
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, types.AuthenticationFailedResponse)
	}
	if err := service.AddRemarkByName(name, userId, req.Remark); err != nil {
		if err == service.ErrNameNotFound {
			return c.JSON(http.StatusFound, types.NameNotFoundResponse)
		}
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.AddRemarkByNameSuccessResponse)
}
