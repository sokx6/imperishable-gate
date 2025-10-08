package links

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service/common"
	linksService "imperishable-gate/internal/server/service/links"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func AddHandler(c echo.Context) error {
	var req request.AddRequest
	// 检测请求是否合法
	if err := c.Bind(&req); err != nil || req.Link == "" {
		logger.Warning("Invalid add link request: empty link")
		return response.InvalidUrlFormatResponse
	}
	// 检测 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		logger.Warning("Invalid URL format: %s", req.Link)
		return response.InvalidUrlFormatResponse
	}
	// 获取用户 ID
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	// 使用 errors.Is 进行错误判断
	err := linksService.AddLink(req.Link, userId)
	if errors.Is(err, common.ErrDatabase) {
		logger.Error("Database error while adding link %s for user %d: %v", req.Link, userId, err)
		return response.DatabaseErrorResponse
	} else if errors.Is(err, common.ErrLinkAlreadyExists) {
		logger.Warning("Link already exists: %s for user %d", req.Link, userId)
		return response.LinkExistsResponse
	} else if err != nil {
		logger.Error("Unknown error while adding link %s for user %d: %v", req.Link, userId, err)
		return response.UnknownErrorResponse
	}

	logger.Success("Link added successfully: %s for user %d", req.Link, userId)
	return c.JSON(http.StatusOK, response.AddLinkSuccessResponse)
}
