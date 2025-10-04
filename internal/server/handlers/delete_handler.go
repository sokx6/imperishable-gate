// handlers/delete.go

package handlers

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

// DeleteHandler 处理通过查询参数删除链接的请求
func DeleteHandler(c echo.Context) error {

	// 从查询参数中获取所有 "link=" 参数值
	links := c.QueryParams()["link"] // 获取同名多个 query 值
	if len(links) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 可选：验证每个 link 是否为合法 URL
	for _, rawLink := range links {
		u, err := url.ParseRequestURI(rawLink)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
			return c.JSON(http.StatusBadRequest, types.InvalidUrlResponse)
		}
	}

	var deletedCount int64

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("url IN ?", links).Delete(&model.Link{})
		if result.Error != nil {
			return result.Error
		}
		deletedCount = result.RowsAffected
		return nil
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	// 返回响应
	if deletedCount == 0 {
		return c.JSON(http.StatusOK, types.LinkNotFoundResponse)
	}

	return c.JSON(http.StatusOK, types.DeleteSuccessResponse)
}
