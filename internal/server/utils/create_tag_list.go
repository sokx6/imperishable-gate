package utils

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func CreateTagList(linktags []string) []model.Tag {
	linktags = removeEmptyAndDuplicate(linktags)
	if len(linktags) == 0 {
		return nil
	}
	var tags []model.Tag
	database.DB.Where("name IN ?", linktags).Find(&tags)
	existing := make(map[string]bool)
	for _, tag := range tags {
		existing[tag.Name] = true
	}

	// 2. 过滤 names，只保留不存在于 existing 中的
	var result []model.Tag
	for _, name := range linktags {
		if !existing[name] {
			result = append(result, model.Tag{Name: name})
		}
	}
	result = append(result, tags...)

	return result
}
