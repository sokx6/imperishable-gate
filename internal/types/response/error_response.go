package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 错误提示消息（前端可通过 HTTP Status 区分错误类型）
var InvalidUrlFormatResponse = echo.NewHTTPError(http.StatusBadRequest, "Invalid URL format")
var InvalidRequestResponse = echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
var DatabaseErrorResponse = echo.NewHTTPError(http.StatusInternalServerError, "Database error")
var RemarkExistsResponse = echo.NewHTTPError(http.StatusConflict, "Remark already exists")
var NameNotFoundResponse = echo.NewHTTPError(http.StatusNotFound, "Name not found")
var NameExistsResponse = echo.NewHTTPError(http.StatusConflict, "Name already exists")
var LinkNotFoundResponse = echo.NewHTTPError(http.StatusNotFound, "Link not found")
var LinkExistsResponse = echo.NewHTTPError(http.StatusConflict, "Link already exists")
var TagNotFoundResponse = echo.NewHTTPError(http.StatusNotFound, "Tag not found")
var UserNameAlreadyExistsResponse = echo.NewHTTPError(http.StatusConflict, "Username already exists")
var EmailAlreadyExistsResponse = echo.NewHTTPError(http.StatusConflict, "Email already registered")
var UserNotFoundResponse = echo.NewHTTPError(http.StatusNotFound, "User not found")
var AuthenticationFailedResponse = echo.NewHTTPError(http.StatusUnauthorized, "Authentication failed")
var InternalServerErrorResponse = echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
var UnknownErrorResponse = echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
