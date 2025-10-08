package handlers

import (
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchByKeywordHandler(c echo.Context) error {
	var resp response.Response
	keyword := c.QueryParam("keyword")
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

	linkList, err := service.ListLinksByKeyword(userId, keyword, page, pageSize)
	if err != nil {
		return response.DatabaseErrorResponse
	}
	resp.Links = linkList
	resp.Message = "Links retrieved successfully"

	return c.JSON(http.StatusOK, resp)
}
