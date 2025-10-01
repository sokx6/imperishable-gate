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
	Action string `json:"action"` // 应为 "delete"
	Link   string `json:"link"`   // 需要删除的链接
}

type DeleteResponse struct {
	Action string `json:"action"` // 回显 "delete"
	Link   string `json:"link"`   // 回显删除的链接
}
