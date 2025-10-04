package handlers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddTagsByLinkHandler(c echo.Context) error {
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "addtagsbylink" || req.Link == "" || req.Tags == nil || len(req.Tags) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.InvalidUrlResponse)
	}

	var link model.Link
	link.Url = req.Link

	if err := database.DB.Where("url = ?", req.Link).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tagList := CreateTagList(req.Tags)
			link.Tags = tagList
			if err := database.DB.Create(&link).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
			}
			return c.JSON(http.StatusOK, types.AddTagsByLinkSuccessResponse)
		} else {
			return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
		}
	}

	tagList := CreateTagList(req.Tags)
	if len(tagList) == 0 {
		return c.JSON(http.StatusOK, types.AddTagsByLinkSuccessResponse)
	}

	if err := database.DB.Model(&link).Association("Tags").Append(tagList); err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.AddTagsByLinkSuccessResponse)

}

func CreateTagList(linktags []string) []model.Tag {
	var tags []model.Tag
	database.DB.Where("name IN ?", linktags).Find(&tags)
	existing := make(map[string]bool)
	for _, tag := range tags {
		existing[tag.Name] = true
	}

	// 2. 过滤 names，只保留不存在于 existing 中的
	var result []model.Tag
	for _, name := range linktags {
		if !existing[name] {
			result = append(result, model.Tag{Name: name})
		}
	}
	result = append(result, tags...)

	return result
}
