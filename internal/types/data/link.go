package data

type Link struct {
	ID          uint     `json:"-"` // 不在 JSON 响应中显示 ID
	Url         string   `json:"url"`
	Tags        []string `json:"tags"`
	Names       []string `json:"names"`
	Remark      string   `json:"remark"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    string   `json:"keywords"`
	StatusCode  int      `json:"status_code"`
	Watching    bool     `json:"watching"`
}
