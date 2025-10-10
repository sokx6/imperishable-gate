// handlers/delete.go

package links

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/response"
)

// DeleteHandler 处理通过查询参数删除链接的请求
func DeleteHandler(c echo.Context) error {

	// 从查询参数中获取所有 "link=" 参数值
	links := c.QueryParams()["link"] // 获取同名多个 query 值
	if len(links) == 0 {
		logger.Warning("Delete failed: no link parameters provided")
		return response.InvalidRequestResponse
	}

	// 验证每个 link 是否为合法 URL
	for _, rawLink := range links {
		u, err := url.ParseRequestURI(rawLink)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
			logger.Warning("Invalid URL format: %s", rawLink)
			return response.InvalidUrlFormatResponse
		}
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		logger.Warning("Delete failed: user not authenticated")
		return response.AuthenticationFailedResponse
	}
	var deletedCount int64

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("url IN ? AND user_id = ?", links, userId).Delete(&model.Link{})
		if result.Error != nil {
			return result.Error
		}
		deletedCount = result.RowsAffected
		return nil
	})

	if err != nil {
		logger.Error("Database error while deleting links for user %d: %v", userId, err)
		return response.DatabaseErrorResponse
	}

	// 返回响应
	if deletedCount == 0 {
		logger.Warning("Delete failed: no matching links found for user %d", userId)
		return response.LinkNotFoundResponse
	}

	return c.JSON(http.StatusOK, response.DeleteSuccessResponse)
}
