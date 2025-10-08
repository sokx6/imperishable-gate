package links

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service/common"
	linksService "imperishable-gate/internal/server/service/links"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func AddHandler(c echo.Context) error {
	var req request.AddRequest
	// 检测请求是否合法
	if err := c.Bind(&req); err != nil || req.Link == "" {
		return response.InvalidUrlFormatResponse
	}
	// 检测 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
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
		return response.DatabaseErrorResponse
	} else if errors.Is(err, common.ErrLinkAlreadyExists) {
		return response.LinkExistsResponse
	} else if err != nil {
		return response.UnknownErrorResponse
	}

	return c.JSON(http.StatusOK, response.AddLinkSuccessResponse)
}
