package links

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/data"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ListByTagHandler(c echo.Context) error {
	tagName := c.Param("tag")
	page, err := utils.GetContentInt(c, "page")
	if err != nil {
		return response.InvalidRequestResponse
	}
	pageSize, err := utils.GetContentInt(c, "page_size")
	if err != nil {
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	// 首先验证 tag 是否存在
	var tag model.Tag
	result := database.DB.First(&tag, "name = ? AND user_id = ?", tagName, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return response.TagNotFoundResponse
		}
		return response.DatabaseErrorResponse
	}

	// 使用 JOIN 查询并在 SQL 层面分页
	var links []model.Link
	err = database.DB.Model(&model.Link{}).
		Joins("INNER JOIN link_tags ON link_tags.link_id = links.id").
		Where("link_tags.tag_id = ? AND links.user_id = ?", tag.ID, userId).
		Order("links.created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Preload("Tags").
		Preload("Names").
		Find(&links).Error
	if err != nil {
		return response.DatabaseErrorResponse
	}

	// 转换为响应格式
	linkList := make([]data.Link, 0, len(links))
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

	return c.JSON(http.StatusOK, response.Response{
		Message: "Links retrieved successfully",
		Links:   linkList,
	})
}
