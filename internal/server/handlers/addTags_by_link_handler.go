package handlers

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func AddTagsByLinkHandler(c echo.Context) error {
	var req request.AddRequest
	if err := c.Bind(&req); err != nil || req.Link == "" || req.Tags == nil || len(req.Tags) == 0 {
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return response.InvalidUrlFormatResponse
	}

	if err := service.AddTagsByLink(req.Link, userId, req.Tags); err != nil {
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.AddTagsByLinkSuccessResponse)

}
