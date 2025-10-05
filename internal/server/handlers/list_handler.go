package handlers

import (
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListHandler(c echo.Context) error {
	var links []model.Link
	userId, ok := utils.GetUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, types.AuthenticationFailedResponse)
	}
	if err := database.DB.Preload("Names").Preload("Tags").Where("user_id = ?", userId).Find(&links).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}
	var linkList []types.Link

	for _, link := range links {
		linkList = append(linkList, types.Link{
			ID:          link.ID,
			Url:         link.Url,
			Tags:        utils.ExtractTagNames(link.Tags),
			Names:       utils.ExtractNames(link.Names),
			Remark:      link.Remark,
			Title:       link.Title,
			Description: link.Description,
			Keywords:    link.Keywords,
		})

	}
	// 返回成功响应
	return c.JSON(http.StatusOK, types.ListResponse{
		Code:    0,
		Message: "Success",
		Data:    linkList,
	})
}
