package utils

import (
	"strings"

	"imperishable-gate/internal/model"
)

func CreateNameList(linknames []string, userId uint) []model.Name {
	var nameList []model.Name
	for _, n := range linknames {
		if trimmed := strings.TrimSpace(n); trimmed != "" {
			nameList = append(nameList, model.Name{Name: trimmed, UserID: userId})
		}
	}
	return nameList
}
