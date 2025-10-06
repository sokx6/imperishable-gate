package handlers

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
)

func AddRemarkHandler(c echo.Context) error {
	var req request.AddRequest
	// 检查请求的有效性
	if err := c.Bind(&req); err != nil || req.Link == "" || req.Remark == "" {
		return response.InvalidRequestResponse
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return response.InvalidUrlFormatResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	if err := service.AddRemarkByLink(req.Link, userId, req.Remark); err != nil {
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.AddRemarkByLinkSuccessResponse)
}
