package response

import "imperishable-gate/internal/types/data"

type ListResponse struct {
	Message string      `json:"message"`
	Data    []data.Link `json:"data"`
}

type ListByNameResponse struct {
	Message string    `json:"message"`
	Data    data.Link `json:"data"`
}
