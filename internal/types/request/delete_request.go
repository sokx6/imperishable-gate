package request

type DeleteRequest struct {
	Url   string   `json:"url"`
	Tags  []string `json:"tags"`
	Names []string `json:"names"`
}
