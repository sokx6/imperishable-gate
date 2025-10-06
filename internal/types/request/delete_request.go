package request

type DeleteRequest struct {
	Tags  []string `json:"tags"`
	Names []string `json:"names"`
}
