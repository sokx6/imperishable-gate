package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddTagsByLinkHandler(c echo.Context) error {
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "addtagsbylink" || req.Link == "" || req.Tags == nil || len(req.Tags) == 0 {
		return c.JSON(400, types.InvalidRequestResponse)
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(400, types.InvalidURLResponse)
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
			return c.JSON(http.StatusOK, types.OKResponse)
		} else {
			return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
		}
	}

	tagList := CreateTagList(req.Tags)
	if len(tagList) == 0 {
		return c.JSON(http.StatusOK, types.OKResponse)
	}

	if err := database.DB.Model(&link).Association("Tags").Append(tagList); err != nil {
		return c.JSON(500, types.NameExistsResponse)
	}

	return c.JSON(200, types.OKResponse)

}

func CreateTagList(linktags []string) []model.Tag {
	var tagList []model.Tag
	for _, n := range linktags {
		if trimmed := strings.TrimSpace(n); trimmed != "" {
			tagList = append(tagList, model.Tag{Name: trimmed})
		}
	}
	return tagList
}
