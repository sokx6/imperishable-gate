package response

import "imperishable-gate/internal/types/data"

type AddResponse struct {
	Response
	Data *data.Link `json:"data,omitempty"`
}
