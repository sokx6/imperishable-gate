// Package handlers 处理 HTTP 请求
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

// AddHandler 处理添加链接请求
func AddHandler(c echo.Context) error {
	var req types.AddRequest

	// 判断请求体数据结构是否与req匹配
	if err := c.Bind(&req); err != nil || req.Action != "add" || req.Link == "" {
		return c.JSON(http.StatusBadRequest, types.AddResponse{
			Code:    -1,
			Message: "Invalid request data",
		})
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.AddResponse{
			Code:    -1,
			Message: "Invalid URL format",
		})
	}

	// 定义link类型
	var link model.Link

	// 开启事务（可选，用于防止并发）
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 先查找是否已存在该 URL
		if err := tx.Where("url = ?", req.Link).First(&link).Error; err == nil {
			// 找到了，说明已存在
			return echo.NewHTTPError(http.StatusConflict, "The link already exists")
		}

		// 创建新记录
		link = model.Link{URL: req.Link}
		if err := tx.Create(&link).Error; err != nil {
			// 如果创建失败（其他原因）
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create link")
		}

		return nil
	})

	// 判断事务返回的错误类型
	if err != nil {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			if httpErr.Code == http.StatusConflict {
				return c.JSON(http.StatusConflict, types.AddResponse{
					Code:    -1,
					Message: "The link already exists",
				})
			}
		}
		return c.JSON(http.StatusInternalServerError, types.AddResponse{
			Code:    -1,
			Message: "Database error",
		})
	}

	// 返回成功结果
	return c.JSON(http.StatusOK, types.AddResponse{
		Code:    0,
		Message: "Added successfully",
		Data: map[string]interface{}{
			"id":  link.ID,
			"url": link.URL,
		},
	})
}
