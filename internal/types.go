package types

import (
	"imperishable-gate/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var InvalidUrlResponse = Response{
	Code:    -1,
	Message: "Invalid URL format",
}

var InvalidRequestResponse = Response{
	Code:    -1,
	Message: "Invalid request",
}

var DatabaseErrorResponse = Response{
	Code:    -1,
	Message: "Database error",
}

var RemarkExistsResponse = Response{
	Code:    -1,
	Message: "Remark already exists",
}

var NameNotFoundResponse = Response{
	Code:    -1,
	Message: "Name not found",
}

var NameExistsResponse = Response{
	Code:    -1,
	Message: "Name already exists",
}

var LinkNotFoundResponse = Response{
	Code:    -1,
	Message: "Link not found",
}

var OKResponse = Response{
	Code:    0,
	Message: "Success",
}

var InvalidUrlFormatResponse = Response{
	Code:    -1,
	Message: "Invalid URL format",
}

var LinkExistsResponse = Response{
	Code:    -1,
	Message: "Link already exists",
}

var AddLinkSuccessResponse = Response{
	Code:    0,
	Message: "Added successfully",
}

var AddNamesSuccessResponse = Response{
	Code:    0,
	Message: "Names added successfully",
}

var AddRemarkByLinkSuccessResponse = Response{
	Code:    0,
	Message: "Remark added successfully",
}

var AddRemarkByNameSuccessResponse = Response{
	Code:    0,
	Message: "Remark added successfully",
}

var AddTagsByLinkSuccessResponse = Response{
	Code:    0,
	Message: "Tags added successfully",
}

var DeleteSuccessResponse = Response{
	Code:    0,
	Message: "Links deleted successfully",
}

var PongResponse = Response{
	Code:    0,
	Message: "pong",
}

var AddTagsByNameSuccessResponse = Response{
	Code:    0,
	Message: "Tags added successfully",
}

var TagNotFoundResponse = Response{
	Code:    -1,
	Message: "Tag not found",
}

var DeleteTagsByNameSuccessResponse = Response{
	Code:    0,
	Message: "Tags deleted successfully",
}

var DeleteNamesByLinkSuccessResponse = Response{
	Code:    0,
	Message: "Names deleted successfully",
}

var DeleteTagsByLinkSuccessResponse = Response{
	Code:    0,
	Message: "Tags deleted successfully",
}

var UserNameAlreadyExistsResponse = Response{
	Code:    -1,
	Message: "Username already exists",
}

var EmailAlreadyExistsResponse = Response{
	Code:    -1,
	Message: "Email already registered",
}

var UserNotFoundResponse = Response{
	Code:    -1,
	Message: "User not found",
}

var AuthenticationFailedResponse = Response{
	Code:    -1,
	Message: "Authentication failed",
}

var RegisterSuccessResponse = Response{
	Code:    0,
	Message: "Registered successfully",
}

var InternalServerErrorResponse = Response{
	Code:    -1,
	Message: "Internal server error",
}

var UnknownErrorResponse = Response{
	Code:    -1,
	Message: "Unknown error",
}

type UserInfo struct {
	UserID   uint
	Username string
}

type PingRequest struct {
	Action  string `json:"action"`  // 应为 "ping"
	Message string `json:"message"` // 客户端发送的消息
}

type PingResponse struct {
	Action  string `json:"action"`  // 回显 "ping"
	Message string `json:"message"` // 固定回复 "pong"
}

type AddRequest struct {
	Action string   `json:"action"` // 应为 "add"
	Link   string   `json:"link"`   // 需要添加的链接
	Names  []string `json:"names"`  // 需要添加的名称列表
	Remark string   `json:"remark"` // 备注，可选
	Tags   []string `json:"tags"`   // 需要添加的标签列表
	Name   string   `json:"name"`   // 链接名称
}

type AddResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`        // 回显添加的链接
	Data    interface{} `json:"data,omitempty"` // 可选，包含新添加链接的详细信息
}

type DeleteRequest struct {
	Action string   `json:"action"` // 应为 "delete"
	Links  []string `json:"link"`   // 需要删除的链接
}

type DeleteResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`        // 删除结果消息
	Data    interface{} `json:"data,omitempty"` // 可选，包含删除链接的详细信息
}

type ListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"` // 列表获取结果消息
	Data    []Link `json:"data"`
}

type AddTagsByLinkRequest struct {
	Action string   `json:"action"` // 应为 "addtagsbylink"
	Link   string   `json:"link"`   // 需要添加标签的链接
	Tags   []string `json:"tags"`   // 需要添加的标签列表
}

type AddTagsResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`        // 回显添加的链接
	Data    interface{} `json:"data,omitempty"` // 可选，包含新添加链接的详细信息
}

type ListByTagRequest struct {
	Action string `json:"action"` // 应为 "listbytag"
	Tag    string `json:"tag"`    // 需要查询的标签
}

type ListByTagResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"` // 列表获取结果消息
	Data    model.Link `json:"data"`    // 可选，包含链接列表
}

type ListByNameResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID          uint
		Url         string
		Tags        []string
		Names       []string
		Remark      string
		Title       string
		Description string
		Keywords    string
	} `json:"data"`
}

type Link struct {
	ID          uint     `json:"id"`
	Url         string   `json:"url"`
	Tags        []string `json:"tags"`
	Names       []string `json:"names"`
	Remark      string   `json:"remark"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    string   `json:"keywords"`
}

type AddRemarkByNameRequest struct {
	Remark string `json:"remark"` // 备注内容
}

type DeleteTagsByNameRequest struct {
	Action string   `json:"action"` // 应为 "deletetagsbyname"
	Url    string   `json:"url"`    // 需要删除标签的链接名称
	Tags   []string `json:"tags"`   // 需要删除的标签列表
}

type DeleteNamesByLinkRequest struct {
	Action string   `json:"action"` // 应为 "deletenamesbylink"
	Url    string   `json:"url"`    // 需要删除名称的链接
	Names  []string `json:"names"`  // 需要删除的名称列表
}

type DeleteTagsByLinkRequest struct {
	Action string   `json:"action"` // 应为 "deletetagsbylink"
	Url    string   `json:"url"`    // 需要删除标签的链接
	Tags   []string `json:"tags"`   // 需要删除的标签列表
}

type UserRegisterRequest struct {
	Username string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
}

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// LoginResult 登录结果结构体，用于返回详细信息
type LoginResult struct {
	Success      bool   `json:"success"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Message      string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
