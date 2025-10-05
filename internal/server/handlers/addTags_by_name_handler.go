package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
)

func AddTagsByNameHandler(c echo.Context) error {
	name := c.Param("name")
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "addtagsbyname" || req.Tags == nil || len(req.Tags) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, types.AuthenticationFailedResponse)
	}
	if err := service.AddTagsByName(name, userId, req.Tags); err != nil {
		if err == service.ErrNameNotFound {
			return c.JSON(http.StatusNotFound, types.NameNotFoundResponse)
		}
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}
	return c.JSON(http.StatusOK, types.AddTagsByNameSuccessResponse)

}
