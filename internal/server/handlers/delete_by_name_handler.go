// handlers/delete.go

package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
)

// DeleteHandler 处理通过查询参数删除链接的请求
func DeleteByNameHandler(c echo.Context) error {

	// 从查询参数中获取所有 "link=" 参数值
	name := c.Param("name") // 获取同名多个 query 值
	if name == "" {
		fmt.Println("Name is empty")
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 可选：验证每个 link 是否为合法 URL

	if id := utils.NameToLinkId(name); id == 0 {
		return c.JSON(http.StatusNotFound, types.NameNotFoundResponse)
	} else if err := database.DB.Delete(&model.Link{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.OKResponse)

}
