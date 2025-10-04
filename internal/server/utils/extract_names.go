package utils

import (
	"imperishable-gate/internal/model"
)

func ExtractNames(Names []model.Name) []string {
	names := make([]string, 0, len(Names)) // 预分配容量，提升性能
	for _, name := range Names {
		names = append(names, name.Name)
	}
	return names
}
