package data

type Link struct {
	ID          uint     `json:"id"`
	Url         string   `json:"url"`
	Tags        []string `json:"tags"`
	Names       []string `json:"names"`
	Remark      string   `json:"remark"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    string   `json:"keywords"`
	StatusCode  int      `json:"status_code"`
}
