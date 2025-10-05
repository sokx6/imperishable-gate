package handlers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
)

var ErrLinkAlreadyExists = errors.New("link already exists")

func AddHandler(c echo.Context) error {
	var req types.AddRequest
	// 检测请求是否合法
	if err := c.Bind(&req); err != nil || req.Action != "add" || req.Link == "" {
		return c.JSON(http.StatusBadRequest, types.InvalidUrlResponse)
	}
	// 检测 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.InvalidUrlFormatResponse)
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, types.AuthenticationFailedResponse)
	}
	switch err := service.AddLink(req.Link, userId); {
	case err == service.ErrDatabase:
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	case err == ErrLinkAlreadyExists:
		return c.JSON(http.StatusConflict, types.LinkExistsResponse)
	case err != nil:
		return c.JSON(http.StatusInternalServerError, types.UnknownErrorResponse)
	default:
		return c.JSON(http.StatusOK, types.AddLinkSuccessResponse)
	}

}
