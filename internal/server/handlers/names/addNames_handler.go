package names

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service/common"
	namesService "imperishable-gate/internal/server/service/names"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

/* var (
	ErrNameAlreadyExists = errors.New("name already exists for another link")
) */

func AddNamesHandler(c echo.Context) error {
	var req request.AddRequest
	if err := c.Bind(&req); err != nil || req.Link == "" || req.Names == nil || len(req.Names) == 0 {
		logger.Warning("Invalid add names request: empty link or names")
		return response.InvalidRequestResponse
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		logger.Warning("Invalid URL format: %s", req.Link)
		return response.InvalidUrlFormatResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := namesService.AddNames(req.Link, userId, req.Names); errors.Is(err, common.ErrNameAlreadyExists) {
		logger.Warning("Name already exists for another link: %v", req.Names)
		return response.NameExistsResponse
	} else if errors.Is(err, common.ErrInvalidRequest) {
		logger.Warning("Invalid add names request: %v", req.Names)
		return response.InvalidRequestResponse
	} else if err != nil {
		logger.Error("Database error while adding names: %v", err)
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.AddNamesSuccessResponse)
}
