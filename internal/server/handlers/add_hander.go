// Package handlers 处理 HTTP 请求
package handlers

import (
	"net/http"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"

	"github.com/labstack/echo/v4"
)

// AddHandler 处理添加链接请求
func AddHandler(c echo.Context) error {
	var req struct {
		Action string `json:"action"` // 应为 "add"
		Link   string `json:"link"`
	}

	// 验证请求
	if err := c.Bind(&req); err != nil || req.Action != "add" || req.Link == "" {
		return c.JSON(http.StatusBadRequest, types.AddResponse{
			Code:    -1,
			Message: "Invalid request data",
		})
	}

	// 插入数据库
	link := model.Link{URL: req.Link}
	result := database.DB.Create(&link)

	// 处理可能的错误
	if result.Error != nil {
		if isUniqueConstraintError(result.Error) {
			return c.JSON(http.StatusConflict, types.AddResponse{
				Code:    -1,
				Message: "The link already exists",
			})
		}
		return c.JSON(http.StatusInternalServerError, types.AddResponse{
			Code:    -1,
			Message: "Database error",
		})
	}

	// 成功响应
	return c.JSON(http.StatusOK, types.AddResponse{
		Code:    0,
		Message: "Added successfully",
		Data: map[string]interface{}{
			"id":  link.ID,
			"url": link.URL,
		},
	})
}

// isUniqueConstraintError 判断是否是唯一性冲突
func isUniqueConstraintError(err error) bool {
	msg := err.Error()
	return msg == `ERROR: duplicate key value violates unique constraint "idx_links_url"` ||
		msg == `ERROR: duplicate key violates unique constraint "links_url_key"`
}
