// handlers/delete.go

package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

// DeleteHandler 处理通过查询参数删除链接的请求
func DeleteTagsByNameHandler(c echo.Context) error {

	var req request.DeleteRequest
	var url = c.QueryParam("url")
	if err := c.Bind(&req); err != nil || url == "" || req.Tags == nil || len(req.Tags) == 0 {
		return response.InvalidRequestResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	if err := service.DeleteTagsByLink(url, userId, req.Tags); err != nil {
		if err == service.ErrLinkNotFound {
			return response.LinkNotFoundResponse
		}
		if err == service.ErrInvalidRequest {
			return response.InvalidRequestResponse
		}
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.DeleteTagsByNameSuccessResponse)

}
