package utils

import (
	"imperishable-gate/internal/model"
)

func ExtractTagNames(tags []model.Tag) []string {
	names := make([]string, 0, len(tags)) // 预分配容量，提升性能
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return names
}
