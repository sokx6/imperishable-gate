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

/* var (
	ErrNameAlreadyExists = errors.New("name already exists for another link")
) */

func AddNamesHandler(c echo.Context) error {
	var req request.AddRequest
	if err := c.Bind(&req); err != nil || req.Link == "" || req.Names == nil || len(req.Names) == 0 {
		return response.InvalidRequestResponse
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return response.InvalidUrlFormatResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := service.AddNames(req.Link, userId, req.Names); errors.Is(err, service.ErrNameAlreadyExists) {
		return response.NameExistsResponse
	} else if errors.Is(err, service.ErrInvalidRequest) {
		return response.InvalidRequestResponse
	} else if err != nil {
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.AddNamesSuccessResponse)
}
