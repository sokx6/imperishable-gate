package handlers

import (
	"fmt"
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"net/http"

	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ListByNameHandler(c echo.Context) error {
	var Name model.Name
	fmt.Println("1")
	// 查询所有记录
	if err :=
		database.DB.
			Preload("Link").
			Preload("Link.Names").
			Preload("Link.Tags").
			Where("name = ?", c.Param("name")).
			First(&Name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, types.NameNotFoundResponse)
		}
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}
	Link := Name.Link

	return c.JSON(http.StatusOK, types.ListByNameResponse{
		Code:    0,
		Message: "Success",
		Data: struct {
			ID          uint
			Url         string
			Tags        []string
			Names       []string
			Remark      string
			Title       string
			Description string
			Keywords    string
		}{
			ID:          Link.ID,
			Url:         Link.Url,
			Tags:        utils.ExtractTagNames(Link.Tags),
			Names:       utils.ExtractNames(Link.Names),
			Remark:      Link.Remark,
			Title:       Link.Title,
			Description: Link.Description,
			Keywords:    Link.Keywords,
		},
	})
}
