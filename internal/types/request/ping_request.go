package request

type PingRequest struct {
	Action  string `json:"action"`  // 必须是 "ping"
	Message string `json:"message"` // 客户端消息（可选）
}
