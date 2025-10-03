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

var ErrLinkAlreadyExists = errors.New("link already exists")

func AddHandler(c echo.Context) error {
	var req types.AddRequest
	if err := c.Bind(&req); err != nil || req.Action != "add" || req.Link == "" {
		return c.JSON(400, map[string]interface{}{
			"code":    -1,
			"message": "Invalid request",
		})
	}

	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(400, map[string]interface{}{
			"code":    -1,
			"message": "Invalid URL format",
		})
	}

	var link model.Link

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("url = ?", req.Link).First(&link)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			link = model.Link{
				Url: req.Link,
			}
			if err := tx.Create(&link).Error; err != nil {
				return err
			}
			return nil
		} else if result.Error != nil {
			return result.Error
		}

		return ErrLinkAlreadyExists

	}); errors.Is(err, ErrLinkAlreadyExists) {
		return c.JSON(400, map[string]interface{}{
			"code":    -1,
			"message": "Link already exists",
		})
	} else if err != nil {
		return c.JSON(500, map[string]interface{}{
			"code":    -1,
			"message": "Database error",
		})
	}

	return c.JSON(200, types.AddResponse{
		Code:    0,
		Message: "Added successfully",
		Data: map[string]interface{}{
			"id":  link.ID,
			"url": link.Url,
		},
	})
}
