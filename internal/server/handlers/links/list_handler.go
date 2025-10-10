package links

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/data"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListHandler(c echo.Context) error {
	var links []model.Link
	page, err := utils.GetContentInt(c, "page")
	if err != nil {
		logger.Warning("Invalid page parameter")
		return response.InvalidRequestResponse
	}
	pageSize, err := utils.GetContentInt(c, "page_size")
	if err != nil {
		logger.Warning("Invalid page_size parameter")
		return response.InvalidRequestResponse
	}
	// 查询记录
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := database.DB.
		Preload("Names").
		Preload("Tags").
		Where("user_id = ?", userId).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&links).Error; err != nil {
		logger.Error("Database error while retrieving links for user %d: %v", userId, err)
		return response.DatabaseErrorResponse
	}
	var linkList []data.Link

	for _, link := range links {
		linkList = append(linkList, data.Link{
			Url:         link.Url,
			Tags:        utils.ExtractTagNames(link.Tags),
			Names:       utils.ExtractNames(link.Names),
			Remark:      link.Remark,
			Title:       link.Title,
			Description: link.Description,
			Keywords:    link.Keywords,
			StatusCode:  link.StatusCode,
			Watching:    link.Watching,
		})
	}
	// 返回成功响应
	logger.Success("Links retrieved successfully for user %d", userId)
	return c.JSON(http.StatusOK, response.Response{
		Message: "Links retrieved successfully",
		Links:   linkList,
	})
}
