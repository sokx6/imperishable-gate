package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

/* var (
	ErrNameAlreadyExists = errors.New("name already exists for another link")
) */

func AddNamesHandler(c echo.Context) error {
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "addname" || req.Link == "" || req.Names == nil || len(req.Names) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.InvalidUrlResponse)
	}

	var link model.Link
	link.Url = req.Link

	if err := database.DB.Where("url = ?", req.Link).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 链接不存在，创建新链接并添加名称
			nameList := CreateNameList(req.Names)
			link.Names = nameList
			if err := database.DB.Create(&link).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
			}
			return c.JSON(http.StatusOK, types.OKResponse)
		} else {
			return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
		}
	}

	// 链接已存在，添加名称
	nameList := CreateNameList(req.Names)
	if len(nameList) == 0 {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	if err := database.DB.Model(&link).Association("Names").Append(nameList); err != nil {
		return c.JSON(http.StatusConflict, types.NameExistsResponse)
	}

	return c.JSON(http.StatusOK, types.AddNamesSuccessResponse)

}

// CreateNameList 创建名称列表，去除空名称
func CreateNameList(linknames []string) []model.Name {
	var nameList []model.Name
	for _, n := range linknames {
		if trimmed := strings.TrimSpace(n); trimmed != "" {
			nameList = append(nameList, model.Name{Name: trimmed})
		}
	}
	return nameList
}
