package handlers

import (
	"errors"
	"net/url"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddNameHandler(c echo.Context) error {
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "addname" || req.Link == "" {
		return c.JSON(400, types.InvalidRequestResponse)
	}

	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(400, types.InvalidURLResponse)
	}

	var link model.Link

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("url = ?", req.Link).First(&link)
		// 如果找不到对应的 Link，创建新的 Link 和关联的 Names
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			link = model.Link{
				Url:   req.Link,
				Names: CreateNameList(req.Names),
			}
			// 创建新的 Link
			if err := tx.Create(&link).Error; err != nil {
				return err
			}
			return nil
		} else if result.Error != nil { // 其他数据库错误
			return result.Error
		}

		return ErrLinkAlreadyExists

	}); errors.Is(err, ErrLinkAlreadyExists) { // 如果 Link 已存在，添加新的 Names
		assocErr := database.DB.Model(&link).
			Association("Names").
			Append(CreateNameList(req.Names))
		if assocErr != nil {
			return c.JSON(500, types.DatabaseErrorResponse)
		}
	} else if err != nil {
		return c.JSON(500, types.DatabaseErrorResponse)
	}
	// 返回成功响应
	return c.JSON(200, types.AddResponse{
		Code:    0,
		Message: "Added successfully",
		Data: map[string]interface{}{
			"id":  link.ID,
			"url": link.Url,
		},
	})
}

func CreateNameList(names []string) []model.Name {
	var nameList []model.Name
	for _, n := range names {
		nameList = append(nameList, model.Name{Name: n})
	}
	return nameList
}
