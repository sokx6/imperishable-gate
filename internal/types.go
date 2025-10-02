package types

type PingRequest struct {
	Action  string `json:"action"`  // 应为 "ping"
	Message string `json:"message"` // 客户端发送的消息
}

type PingResponse struct {
	Action  string `json:"action"`  // 回显 "ping"
	Message string `json:"message"` // 固定回复 "pong"
}

type AddRequest struct {
	Action string `json:"action"` // 应为 "add"
	Link   string `json:"link"`   // 需要添加的链接
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
	Code    int      `json:"code"`
	Message string   `json:"message"` // 列表获取结果消息
	Data    []string `json:"data"`    // 可选，包含链接列表
}

type AddTagsRequest struct {
	Action string   `json:"action"` // 应为 "addtags"
	Link   string   `json:"link"`   // 需要添加标签的链接
	Tags   []string `json:"tags"`   // 需要添加的标签列表
}

type AddTagsResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`        // 回显添加的链接
	Data    interface{} `json:"data,omitempty"` // 可选，包含新添加链接的详细信息
}
