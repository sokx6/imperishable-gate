package handlers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

var ErrLinkAlreadyExists = errors.New("link already exists")

func AddHandler(c echo.Context) error {
	var req request.AddRequest
	// 检测请求是否合法
	if err := c.Bind(&req); err != nil || req.Link == "" {
		return response.InvalidUrlFormatResponse
	}
	// 检测 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return response.InvalidUrlFormatResponse
	}
	// 获取用户 ID
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	switch err := service.AddLink(req.Link, userId); {
	case err == service.ErrDatabase:
		return response.DatabaseErrorResponse
	case err == ErrLinkAlreadyExists:
		return response.LinkExistsResponse
	case err != nil:
		return response.UnknownErrorResponse
	default:
		return c.JSON(http.StatusOK, response.AddLinkSuccessResponse)
	}

}
