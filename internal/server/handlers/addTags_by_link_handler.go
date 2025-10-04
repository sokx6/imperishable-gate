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
	"imperishable-gate/internal/server/utils"
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
			tagList := utils.CreateTagList(req.Tags)
			link.Tags = tagList
			if err := database.DB.Create(&link).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
			}
			return c.JSON(http.StatusOK, types.AddTagsByLinkSuccessResponse)
		} else {
			return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
		}
	}

	tagList := utils.CreateTagList(req.Tags)
	if len(tagList) == 0 {
		return c.JSON(http.StatusOK, types.AddTagsByLinkSuccessResponse)
	}

	if err := database.DB.Model(&link).Association("Tags").Append(tagList); err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.AddTagsByLinkSuccessResponse)

}
