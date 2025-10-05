package handlers

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
)

func AddRemarkHandler(c echo.Context) error {
	var req types.AddRequest
	// 检查请求的有效性
	if err := c.Bind(&req); err != nil || req.Action != "addremark" || req.Link == "" || req.Remark == "" {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.InvalidUrlResponse)
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, types.AuthenticationFailedResponse)
	}

	if err := service.AddRemarkByLink(req.Link, userId, req.Remark); err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.AddRemarkByLinkSuccessResponse)
}
