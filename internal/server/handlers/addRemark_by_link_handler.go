package handlers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddRemarkHandler(c echo.Context) error {
	var req types.AddRequest
	// 检查请求的有效性
	if err := c.Bind(&req); err != nil || req.Action != "addremark" || req.Link == "" || req.Remark == "" {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 验证 URL 格式
	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(http.StatusBadRequest, types.InvalidUrlResponse)
	}

	var link model.Link

	result := database.DB.Where("url = ?", req.Link).First(&link)
	// 如果找不到对应的 Link，创建新的 Link 和关联的 Names
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		link = model.Link{
			Url:    req.Link,
			Remark: req.Remark,
		}
		// 创建新的 Link
		if err := database.DB.Create(&link).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
		}
	} else if result.Error != nil { // 其他数据库错误
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	if err := database.DB.Model(link).Update("Remark", req.Remark).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.AddRemarkByLinkSuccessResponse)
}
