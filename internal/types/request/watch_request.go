package request

type WatchByUrlRequest struct {
	Url   string `json:"url"`
	Watch bool   `json:"watch"`
}
type WatchByNameRequest struct {
	Name  string `json:"name"`
	Watch bool   `json:"watch"`
}
