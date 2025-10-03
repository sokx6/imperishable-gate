package handlers

import (
	"fmt"
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"net/http"

	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ListByNameHandler(c echo.Context) error {
	var Name model.Name
	fmt.Println("1")
	// 查询所有记录
	if err := database.DB.Preload("Link").Preload("Link.Names").Where("name = ?", c.Param("name")).First(&Name).Error; err != nil {
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
			ID     uint
			Url    string
			Tags   []string
			Names  []string
			Remark string
		}{
			ID:     Link.ID,
			Url:    Link.Url,
			Tags:   ExtractTagNames(Link.Tags),
			Names:  ExtractNames(Link.Names),
			Remark: Link.Remark,
		},
	})
}

func ExtractTagNames(tags []model.Tag) []string {
	names := make([]string, 0, len(tags)) // 预分配容量，提升性能
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return names
}

func ExtractNames(Names []model.Name) []string {
	names := make([]string, 0, len(Names)) // 预分配容量，提升性能
	for _, name := range Names {
		names = append(names, name.Name)
	}
	return names
}
