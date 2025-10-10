package tags

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	tagsService "imperishable-gate/internal/server/service/tags"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func AddTagsByLinkHandler(c echo.Context) error {
	var req request.AddRequest
	if err := c.Bind(&req); err != nil || req.Link == "" || req.Tags == nil || len(req.Tags) == 0 {
		logger.Warning("Invalid add tags request: empty link or tags")
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		logger.Warning("Invalid URL format: %s", req.Link)
		return response.InvalidUrlFormatResponse
	}

	if err := tagsService.AddTagsByLink(req.Link, userId, req.Tags); err != nil {
		logger.Error("Database error while adding tags: %v", err)
		return response.DatabaseErrorResponse
	}

	logger.Success("Tags added successfully for link %s by user %d", req.Link, userId)
	return c.JSON(http.StatusOK, response.AddTagsByLinkSuccessResponse)

}
