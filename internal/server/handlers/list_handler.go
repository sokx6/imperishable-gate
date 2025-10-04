package handlers

import (
	"fmt"
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListHandler(c echo.Context) error {
	var links []model.Link
	fmt.Println("1")
	// 查询所有记录
	if err := database.DB.Preload("Names").Preload("Tags").Find(&links).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, types.ListResponse{
			Code:    -1,
			Message: "Failed to retrieve links",
		})
	}
	var linkList []types.Link

	for _, link := range links {
		linkList = append(linkList, types.Link{
			ID:     link.ID,
			Url:    link.Url,
			Tags:   utils.ExtractTagNames(link.Tags),
			Names:  utils.ExtractNames(link.Names),
			Remark: link.Remark,
		})

	}
	fmt.Println(linkList)

	// 返回成功响应
	return c.JSON(http.StatusOK, types.ListResponse{
		Code:    0,
		Message: "Success",
		Data:    linkList,
	})
}
