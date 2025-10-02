package handlers

import (
	"errors"
	"fmt"
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// AddTags 处理添加标签请求
func AddTags(c echo.Context) error {
	var req types.AddTagsRequest

	// 绑定并验证请求
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

	// 查找链接是否存在，并加锁避免并发修改（可选 FOR UPDATE）
	var link model.Link

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 使用 SELECT ... FOR UPDATE 锁住行（可选）
		result := tx.Where("url = ?", req.Link).First(&link)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("link not found") // 自定义错误用于判断
		}
		if result.Error != nil {
			return result.Error
		}

		// 合并标签并去重
		mergedTags := removeDuplicates(append(link.Tags, req.Tags...))
		link.Tags = mergedTags

		return tx.Model(&link).Update("tags", pq.Array(mergedTags)).Error
	})

	// 处理事务错误
	if err != nil {
		switch {
		case err.Error() == "link not found":
			return c.JSON(http.StatusNotFound, types.AddResponse{
				Code:    -1,
				Message: "The link does not exist",
			})
		default:
			// 日志记录 err
			return c.JSON(http.StatusInternalServerError, types.AddResponse{
				Code:    -1,
				Message: "Database error: " + err.Error(),
			})
		}
	}

	// 成功响应
	return c.JSON(http.StatusOK, types.AddResponse{
		Code:    0,
		Message: "Added successfully",
		Data: map[string]interface{}{
			"id":   link.Id,
			"url":  link.Url,
			"tags": link.Tags, // 可选返回最新 tags
		},
	})
}

// 工具函数：去除字符串 slice 中的重复项
func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, item := range slice {
		if !seen[item] && item != "" { // 忽略空字符串
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
