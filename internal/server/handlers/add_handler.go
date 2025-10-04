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
		return c.JSON(400, types.InvalidUrlResponse)
	}

	if _, err := url.ParseRequestURI(req.Link); err != nil {
		return c.JSON(400, types.InvalidUrlFormatResponse)
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
		return c.JSON(400, types.LinkExistsResponse)
	} else if err != nil {
		return c.JSON(500, types.DatabaseErrorResponse)
	}

	return c.JSON(200, types.AddLinkSuccessResponse)
}
