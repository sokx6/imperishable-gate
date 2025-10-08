// handlers/delete.go

package links

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/response"
)

// DeleteHandler 处理通过查询参数删除链接的请求
func DeleteByNameHandler(c echo.Context) error {

	name := c.Param("name")
	if name == "" {
		logger.Warning("Delete by name failed: name parameter is empty")
		return response.InvalidRequestResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	if id := utils.GetLinkIDByName(name, userId); id == 0 {
		logger.Warning("Delete by name failed: name '%s' not found for user %d", name, userId)
		return response.NameNotFoundResponse
	} else if err := database.DB.Delete(&model.Link{}, id).Error; err != nil {
		logger.Error("Database error while deleting link by name '%s': %v", name, err)
		return response.DatabaseErrorResponse
	}

	logger.Success("Link with name '%s' deleted successfully by user %d", name, userId)
	return c.JSON(http.StatusOK, response.DeleteSuccessResponse)

}
