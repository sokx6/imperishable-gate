package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddRemarkByNameHandler(c echo.Context) error {
	var req types.AddRemarkByNameRequest
	name := c.Param("name")
	// 检查请求的有效性
	if err := c.Bind(&req); err != nil || name == "" || req.Remark == "" {
		return c.JSON(400, types.InvalidRequestResponse)
	}

	var Name model.Name
	if err := database.DB.Where("name = ?", name).Take(&Name).Error; err != nil {
		return c.JSON(http.StatusNotFound, types.LinkNotFoundResponse)
	}
	if err := database.DB.Model(&model.Link{ID: Name.LinkID}).Update("Remark", req.Remark).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(200, types.AddRemarkByNameSuccessResponse)
}
