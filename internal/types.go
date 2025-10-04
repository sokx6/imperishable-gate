package types

import "imperishable-gate/internal/model"

var InvalidUrlResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Invalid URL format",
}

var InvalidRequestResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Invalid request",
}

var DatabaseErrorResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Database error",
}

var RemarkExistsResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Remark already exists",
}

var NameNotFoundResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Name not found",
}

var NameExistsResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Name already exists",
}

var LinkNotFoundResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Link not found",
}

var OKResponse = struct {
	Code    int
	Message string
}{
	Code:    0,
	Message: "Success",
}

var InvalidUrlFormatResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Invalid URL format",
}

var LinkExistsResponse = struct {
	Code    int
	Message string
}{
	Code:    -1,
	Message: "Link already exists",
}

var AddLinkSuccessResponse = struct {
	Code    int
	Message string
}{
	Code:    0,
	Message: "Added successfully",
}

var AddNamesSuccessResponse = struct {
	Code    int
	Message string
}{
	Code:    0,
	Message: "Names added successfully",
}

var AddRemarkByLinkSuccessResponse = struct {
	Code    int
	Message string
}{
	Code:    0,
	Message: "Remark added successfully",
}

var AddRemarkByNameSuccessResponse = struct {
	Code    int
	Message string
}{
	Code:    0,
	Message: "Remark added successfully",
}

var AddTagsByLinkSuccessResponse = struct {
	Code    int
	Message string
}{
	Code:    0,
	Message: "Tags added successfully",
}

var DeleteSuccessResponse = struct {
	Code    int
	Message string
}{
	Code:    0,
	Message: "Links deleted successfully",
}

var PongResponse = struct {
	Code    int
	Message string
}{
	Code:    0,
	Message: "pong",
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
		ID     uint
		Url    string
		Tags   []string
		Names  []string
		Remark string
	} `json:"data"`
}

type Link struct {
	ID     uint     `json:"id"`
	Url    string   `json:"url"`
	Tags   []string `json:"tags"`
	Names  []string `json:"names"`
	Remark string   `json:"remark"`
}

type AddRemarkByNameRequest struct {
	Remark string `json:"remark"` // 备注内容
}
