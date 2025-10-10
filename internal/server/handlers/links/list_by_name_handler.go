package links

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/data"
	"imperishable-gate/internal/types/response"
	"net/http"

	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ListByNameHandler(c echo.Context) error {
	var Name model.Name
	// 查询所有记录
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err :=
		database.DB.
			Preload("Link").
			Preload("Link.Names").
			Preload("Link.Tags").
			Where("name = ? AND user_id = ?", c.Param("name"), userId).
			First(&Name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warning("Name not found: %s for user %d", c.Param("name"), userId)
			return response.NameNotFoundResponse
		}
		logger.Error("Database error while retrieving name %s for user %d: %v", c.Param("name"), userId, err)
		return response.DatabaseErrorResponse
	}
	Link := Name.Link

	responseLink := data.Link{
		Url:         Link.Url,
		Tags:        utils.ExtractTagNames(Link.Tags),
		Names:       utils.ExtractNames(Link.Names),
		Remark:      Link.Remark,
		Title:       Link.Title,
		Description: Link.Description,
		Keywords:    Link.Keywords,
		StatusCode:  Link.StatusCode,
		Watching:    Link.Watching,
	}

	logger.Success("Link with name '%s' retrieved successfully for user %d", c.Param("name"), userId)
	return c.JSON(http.StatusOK, response.Response{
		Message: "Link retrieved successfully",
		Links:   []data.Link{responseLink},
	})
}
