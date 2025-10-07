// handlers/delete.go

package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func DeleteTagsByNameHandler(c echo.Context) error {

	var req request.DeleteRequest
	var name = c.Param("name")
	if err := c.Bind(&req); err != nil || req.Tags == nil || len(req.Tags) == 0 {
		return response.InvalidRequestResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	if err := service.DeleteTagsByName(name, userId, req.Tags); err != nil {
		if errors.Is(err, service.ErrLinkNotFound) {
			return response.LinkNotFoundResponse
		}
		if errors.Is(err, service.ErrInvalidRequest) {
			return response.InvalidRequestResponse
		}
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.DeleteTagsByNameSuccessResponse)

}
