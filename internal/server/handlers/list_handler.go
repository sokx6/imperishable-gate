package handlers

import (
	"fmt"
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListHandler(c echo.Context) error {
	var links []model.Link
	fmt.Println("1")
	// 查询所有记录
	if err := database.DB.Find(&links).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, types.ListResponse{
			Code:    -1,
			Message: "Failed to retrieve links",
		})
	}

	// 提取仅包含 URL 的字符串切片（可选）
	urls := make([]string, len(links))
	for i, link := range links {
		urls[i] = link.Url
	}

	// 返回成功响应
	return c.JSON(http.StatusOK, types.ListResponse{
		Code:    0,
		Message: "Success",
		Data:    urls, // 或返回完整对象数组：links
	})
}
