package types

type PingRequest struct {
	Action  string `json:"action"`  // 应为 "ping"
	Message string `json:"message"` // 客户端发送的消息
}

type PingResponse struct {
	Action  string `json:"action"`  // 回显 "ping"
	Message string `json:"message"` // 固定回复 "pong"
}
