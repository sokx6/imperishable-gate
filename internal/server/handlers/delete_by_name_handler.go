// handlers/delete.go

package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/response"
)

// DeleteHandler 处理通过查询参数删除链接的请求
func DeleteByNameHandler(c echo.Context) error {

	// 从查询参数中获取所有 "link=" 参数值
	name := c.Param("name") // 获取同名多个 query 值
	if name == "" {
		fmt.Println("Name is empty")
		return response.InvalidRequestResponse
	}

	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	if id := utils.NameToLinkId(name, userId); id == 0 {
		return response.NameNotFoundResponse
	} else if err := database.DB.Delete(&model.Link{}, id).Error; err != nil {
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.DeleteSuccessResponse)

}
