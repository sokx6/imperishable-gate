package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetContentInt(c echo.Context, param string) (int, error) {
	contentStr := c.QueryParam(param)
	content, err := strconv.Atoi(contentStr)
	if err != nil {
		return 0, err
	}
	return content, nil
}
