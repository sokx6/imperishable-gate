package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/server/service"
)

func DeleteNamesByLinkHandler(c echo.Context) error {
	var req types.DeleteNamesByLinkRequest
	url := c.QueryParam("url")
	if err := c.Bind(&req); err != nil || req.Action != "deletenamesbylink" || url == "" || req.Names == nil || len(req.Names) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}
	if err := service.DeleteNamesByLink(url, req.Names); err != nil {
		if err == service.ErrLinkNotFound {
			return c.JSON(http.StatusFound, types.LinkNotFoundResponse)
		}
		if err == service.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
		}
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.DeleteNamesByLinkSuccessResponse)
}
