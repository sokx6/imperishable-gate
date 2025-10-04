package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddTagsByNameHandler(c echo.Context) error {
	name := c.Param("name")
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "addtagsbyname" || req.Tags == nil || len(req.Tags) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	tagList := CreateTagList(req.Tags)

	var Name model.Name
	if err := database.DB.Where("name = ?", name).Take(&Name).Error; err != nil {
		return c.JSON(http.StatusNotFound, types.LinkNotFoundResponse)
	}
	if err := database.DB.Model(&model.Link{ID: Name.LinkID}).Update("Tags", tagList).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}
	return c.JSON(http.StatusOK, types.AddTagsByNameSuccessResponse)
}
