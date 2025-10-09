package links

import (
	linksService "imperishable-gate/internal/server/service/links"
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
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	if err != nil {
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}

	linkList, err := linksService.ListLinksByKeyword(userId, keyword, page, pageSize)
	if err != nil {
		return response.DatabaseErrorResponse
	}
	resp.Links = linkList
	resp.Message = "Links retrieved successfully"

	return c.JSON(http.StatusOK, resp)
}
