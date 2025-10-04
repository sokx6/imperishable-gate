// handlers/delete.go

package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/server/service"
)

// DeleteHandler 处理通过查询参数删除链接的请求
func DeleteTagsByNameHandler(c echo.Context) error {

	var req types.DeleteTagsByNameRequest
	var url = c.QueryParam("url")
	if err := c.Bind(&req); err != nil || req.Action != "deletetagsbyname" || url == "" || req.Tags == nil || len(req.Tags) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}
	if err := service.DeleteTagsByLink(url, req.Tags); err != nil {
		if err == service.ErrLinkNotFound {
			return c.JSON(http.StatusFound, types.LinkNotFoundResponse)
		}
		if err == service.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
		}
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.DeleteTagsByNameSuccessResponse)

}
