package handlers

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/server/service"
)

func AddTagsByLinkHandler(c echo.Context) error {
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "addtagsbylink" || req.Link == "" || req.Tags == nil || len(req.Tags) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.InvalidUrlResponse)
	}

	if err := service.AddTagsByLink(req.Link, req.Tags); err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.AddTagsByLinkSuccessResponse)

}
