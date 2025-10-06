package handlers

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/data"
	"imperishable-gate/internal/types/response"
	"net/http"

	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ListByNameHandler(c echo.Context) error {
	var Name model.Name
	// 查询所有记录
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err :=
		database.DB.
			Preload("Link").
			Preload("Link.Names").
			Preload("Link.Tags").
			Where("name = ? AND user_id = ?", c.Param("name"), userId).
			First(&Name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NameNotFoundResponse
		}
		return response.DatabaseErrorResponse
	}
	Link := Name.Link

	responseLink := data.Link{
		ID:          Link.ID,
		Url:         Link.Url,
		Tags:        utils.ExtractTagNames(Link.Tags),
		Names:       utils.ExtractNames(Link.Names),
		Remark:      Link.Remark,
		Title:       Link.Title,
		Description: Link.Description,
		Keywords:    Link.Keywords,
		StatusCode:  Link.StatusCode,
	}
	return c.JSON(http.StatusOK, response.ListByNameResponse{
		Message: "Link retrieved successfully",
		Data:    responseLink,
	})
}
