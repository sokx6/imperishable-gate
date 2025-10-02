// handlers/upsert_link_handler.go
package handlers

import (
	"errors"
	"net/url"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddWithTagsHandler(c echo.Context) error {
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "add" || req.Link == "" {
		return c.JSON(400, map[string]interface{}{
			"code":    -1,
			"message": "Invalid request",
		})
	}

	_, err := url.ParseRequestURI(req.Link)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"code":    -1,
			"message": "Invalid URL format",
		})
	}

	tagNames := removeDuplicates(req.Tags)
	var link model.Link

	// Step 1: 查看是否已有该链接
	result := database.DB.Preload("Tags").Where("url = ?", req.Link).First(&link)
	isCreating := false

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		link.Url = req.Link
		link.Note = ""
		isCreating = true
	} else if result.Error != nil {
		return c.JSON(500, map[string]interface{}{
			"code":    -1,
			"message": "Database error",
		})
	}

	// Step 2: 准备标签列表（关键！要先查出来已有的）
	var finalTags []model.Tag

	if len(tagNames) > 0 {
		// 查找所有已经存在的 tag
		var existingTags []model.Tag
		database.DB.Where("name IN ?", tagNames).Find(&existingTags)

		// 建立 name -> Tag 映射（包含真实 ID）
		existingNameToTag := make(map[string]model.Tag)
		for _, t := range existingTags {
			existingNameToTag[t.Name] = t
		}

		// 构造最终需要关联的 Tag 列表
		for _, name := range tagNames {
			if name == "" {
				continue
			}
			if existing, ok := existingNameToTag[name]; ok {
				// 复用已有 tag（带 ID）
				finalTags = append(finalTags, existing)
			} else {
				// 新标签：仅声明名字，GORM 会在 Save 时自动插入
				finalTags = append(finalTags, model.Tag{Name: name})
			}
		}
	}

	// Step 3: 关联到链接
	link.Tags = finalTags

	// Step 4: 保存（启用 FullSaveAssociations 保证正确级联）
	err = database.DB.Session(&gorm.Session{
		FullSaveAssociations: true,
	}).Save(&link).Error

	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"code":    -1,
			"message": "Save failed: " + err.Error(),
		})
	}

	message := "Updated successfully"
	if isCreating {
		message = "Added successfully"
	}

	return c.JSON(200, types.AddResponse{
		Code:    0,
		Message: message,
		Data: map[string]interface{}{
			"url":  link.Url,
			"tags": link.Tags,
			"note": link.Note,
		},
	})
}

func removeDuplicates(input []string) []string {
	seen := make(map[string]struct{})
	var result []string

	for _, item := range input {
		if _, ok := seen[item]; !ok && item != "" {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}
