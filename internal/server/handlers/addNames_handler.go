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

/* var (
	ErrNameAlreadyExists = errors.New("name already exists for another link")
) */

func AddNamesHandler(c echo.Context) error {
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "addname" || req.Link == "" || req.Names == nil || len(req.Names) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.InvalidUrlResponse)
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, types.AuthenticationFailedResponse)
	}
	if err := service.AddNames(req.Link, userId, req.Names); errors.Is(err, service.ErrNameAlreadyExists) {
		return c.JSON(http.StatusBadRequest, types.NameExistsResponse)
	} else if errors.Is(err, service.ErrInvalidRequest) {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.AddNamesSuccessResponse)

}
