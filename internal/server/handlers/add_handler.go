// handlers/upsert_link_handler.go
package handlers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"gorm.io/gorm"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddWithTagsHandler(c echo.Context) error {
	var req types.AddRequest

	// 绑定并校验基本数据
	if err := c.Bind(&req); err != nil || req.Action != "add" || req.Link == "" {
		return c.JSON(http.StatusBadRequest, types.AddResponse{
			Code:    -1,
			Message: "Invalid request data",
		})
	}

	// 校验 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.AddResponse{
			Code:    -1,
			Message: "Invalid URL format",
		})
	}

	var link model.Link
	var isCreated bool

	// 使用事务处理原子操作
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("url = ?", req.Link).First(&link)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 链接不存在 → 创建新记录，并带上去重后的标签
			tags := removeDuplicates(req.Tags)
			link = model.Link{
				Url: req.Link,
			}
			if err := tx.Create(&link).Error; err != nil {
				return err
			}
			if len(tags) > 0 {
				if err := tx.Model(&link).Update("tags", pq.Array(tags)).Error; err != nil {
					return err
				}
				link.Tags = tags // 手动同步内存中的值用于返回
			}
			isCreated = true
		} else if result.Error != nil {
			// 其他数据库错误
			return result.Error
		} else {
			// 链接已存在 → 合并新旧标签并更新
			mergedTags := removeDuplicates(append(link.Tags, req.Tags...))
			if len(mergedTags) > 0 {
				link.Tags = mergedTags
				if err := tx.Model(&link).Update("tags", pq.Array(mergedTags)).Error; err != nil {
					return err
				}
			}
			isCreated = false
		}

		return nil
	})

	// 处理事务错误
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.AddResponse{
			Code:    -1,
			Message: "Database error: " + err.Error(),
		})
	}

	// 返回成功响应
	message := "Updated successfully"
	if isCreated {
		message = "Added successfully"
	}

	return c.JSON(http.StatusOK, types.AddResponse{
		Code:    0,
		Message: message,
		Data: map[string]interface{}{
			"id":   link.Id,
			"url":  link.Url,
			"tags": link.Tags,
		},
	})
}

// 工具函数：去重
func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, item := range slice {
		if item != "" && !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
