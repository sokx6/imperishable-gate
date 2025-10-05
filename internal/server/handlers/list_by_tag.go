package handlers

import (
	"errors"
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ListByTagHandler(c echo.Context) error {
	tagName := c.Param("tag")
	var tag model.Tag
	userId, ok := utils.GetUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, types.AuthenticationFailedResponse)
	}
	result := database.DB.Preload("Links.Tags").Preload("Links.Names").First(&tag, "name = ? AND user_id = ?", tagName, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, types.TagNotFoundResponse)
		}
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}
	linkList := make([]types.Link, 0)
	for _, link := range tag.Links {
		linkList = append(linkList, types.Link{
			ID:          link.ID,
			Url:         link.Url,
			Tags:        utils.ExtractTagNames(link.Tags),
			Names:       utils.ExtractNames(link.Names),
			Remark:      link.Remark,
			Title:       link.Title,
			Description: link.Description,
			Keywords:    link.Keywords,
		})
	}
	return c.JSON(http.StatusOK, linkList)
}
