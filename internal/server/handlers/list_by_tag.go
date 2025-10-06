package handlers

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/data"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ListByTagHandler(c echo.Context) error {
	tagName := c.Param("tag")
	var tag model.Tag
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	result := database.DB.Preload("Links.Tags").Preload("Links.Names").First(&tag, "name = ? AND user_id = ?", tagName, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return response.TagNotFoundResponse
		}
		return response.DatabaseErrorResponse
	}
	linkList := make([]data.Link, 0)
	for _, link := range tag.Links {
		linkList = append(linkList, data.Link{
			ID:          link.ID,
			Url:         link.Url,
			Tags:        utils.ExtractTagNames(link.Tags),
			Names:       utils.ExtractNames(link.Names),
			Remark:      link.Remark,
			Title:       link.Title,
			Description: link.Description,
			Keywords:    link.Keywords,
			StatusCode:  link.StatusCode,
		})
	}
	return c.JSON(http.StatusOK, response.ListResponse{
		Message: "Links retrieved successfully",
		Data:    linkList,
	})
}
