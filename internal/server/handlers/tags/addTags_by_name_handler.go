package tags

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service/common"
	tagsService "imperishable-gate/internal/server/service/tags"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func AddTagsByNameHandler(c echo.Context) error {
	name := c.Param("name")
	var req request.AddRequest
	if err := c.Bind(&req); err != nil || req.Tags == nil || len(req.Tags) == 0 {
		logger.Warning("Invalid add tags request: empty tags")
		return response.InvalidRequestResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := tagsService.AddTagsByName(name, userId, req.Tags); err != nil {
		if errors.Is(err, common.ErrNameNotFound) {
			logger.Warning("Name not found: %s", name)
			return response.NameNotFoundResponse
		}
		logger.Error("Database error while adding tags: %v", err)
		return response.DatabaseErrorResponse
	}

	logger.Success("Tags added successfully for name %s by user %d", name, userId)
	return c.JSON(http.StatusOK, response.AddTagsByNameSuccessResponse)

}
