package auth

import (
	"net/http"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"github.com/labstack/echo/v4"
)

func LogoutHandler(c echo.Context) error {
	var req request.LogoutRequest
	if err := c.Bind(&req); err != nil {
		logger.Warning("Invalid logout request: %v", err)
		return response.InvalidRequestResponse
	}

	// 获取当前已认证用户的 ID
	userId, ok := utils.GetUserID(c)
	if !ok {
		logger.Warning("Failed to get user ID")
		return response.InternalServerErrorResponse
	}

	var tokenRecord model.RefreshToken
	if err := database.DB.Where("token = ?", req.RefreshToken).First(&tokenRecord).Error; err != nil {
		logger.Warning("Failed to find token: %v", err)
		return response.InvalidRequestResponse
	}

	// 验证 token 是否属于当前用户
	if tokenRecord.UserID != userId {
		logger.Warning("Token does not belong to the authenticated user")
		return response.AuthenticationFailedResponse
	}

	tokenRecord.Revoked = true
	database.DB.Save(&tokenRecord)

	logger.Success("User %d logged out successfully", userId)
	return c.JSON(http.StatusOK, response.Response{
		Message: "Logged out successfully",
	})
}
