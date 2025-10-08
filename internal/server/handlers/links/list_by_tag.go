package links

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
	page, err := utils.GetContentInt(c, "page")
	if err != nil {
		return response.InvalidRequestResponse
	}
	pageSize, err := utils.GetContentInt(c, "page_size")
	if err != nil {
		return response.InvalidRequestResponse
	}
	var tag model.Tag
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	result := database.DB.
		Preload("Links.Tags").
		Preload("Links.Names").
		First(&tag, "name = ? AND user_id = ?", tagName, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return response.TagNotFoundResponse
		}
		return response.DatabaseErrorResponse
	}
	total := int64(len(tag.Links))
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > int(total) {
		start = int(total)
	}
	if end > int(total) {
		end = int(total)
	}
	linkList := make([]data.Link, 0)
	for i := start; i < end; i++ {
		link := tag.Links[i]
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
	return c.JSON(http.StatusOK, response.Response{
		Message: "Links retrieved successfully",
		Links:   linkList,
	})
}
