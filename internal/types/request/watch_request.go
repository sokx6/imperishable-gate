package request

type WatchRequest struct {
	Url   string `json:"url"`
	Watch bool   `json:"watch"`
}
