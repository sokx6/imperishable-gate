package remarks

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/service/common"
	remarksService "imperishable-gate/internal/server/service/remarks"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func AddRemarkByNameHandler(c echo.Context) error {
	var req request.AddRequest
	name := c.Param("name")
	// 检查请求的有效性
	if err := c.Bind(&req); err != nil || name == "" || req.Remark == "" {
		logger.Warning("Invalid add remark request: empty name or remark")
		return response.InvalidRequestResponse
	}
	userId, ok := utils.GetUserID(c)
	if !ok {
		return response.AuthenticationFailedResponse
	}
	if err := remarksService.AddRemarkByName(name, userId, req.Remark); err != nil {
		if errors.Is(err, common.ErrNameNotFound) {
			logger.Warning("Name not found: %s", name)
			return response.NameNotFoundResponse
		}
		logger.Error("Database error while adding remark: %v", err)
		return response.DatabaseErrorResponse
	}

	logger.Success("Remark added successfully for name %s by user %d", name, userId)
	return c.JSON(http.StatusOK, response.AddRemarkByNameSuccessResponse)
}
