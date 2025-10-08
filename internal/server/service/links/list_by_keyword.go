package links

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/types/data"
	"strings"
)

func ListLinksByKeyword(userID uint, keyword string, page int, pageSize int) ([]data.Link, error) {
	var links []model.Link

	// 防止空 keyword 导致全表扫描
	if strings.TrimSpace(keyword) == "" {
		return []data.Link{}, nil
	}

	searchPattern := "%" + keyword + "%"

	// 查询符合条件的数据并分页
	err := database.DB.Model(&model.Link{}).
		Joins("LEFT JOIN link_tags ON link_tags.link_id = links.id").
		Joins("LEFT JOIN tags ON tags.id = link_tags.tag_id AND tags.user_id = ?", userID).
		Joins("LEFT JOIN names ON names.link_id = links.id AND names.user_id = ?", userID).
		Where("links.user_id = ?", userID).
		Where("links.url LIKE ? OR links.remark LIKE ? OR links.title LIKE ? OR "+
			"links.description LIKE ? OR links.keywords LIKE ? OR "+
			"tags.name LIKE ? OR names.name LIKE ?",
			searchPattern, searchPattern,
			searchPattern, searchPattern,
			searchPattern, searchPattern,
			searchPattern).
		Group("links.id").
		Order("links.created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Preload("Tags").  // 预加载完整 Tags
		Preload("Names"). // 预加载完整 Names
		Find(&links).Error
	if err != nil {
		return nil, err
	}

	// 转换为 data.Link 切片
	var linkList []data.Link
	for _, link := range links {
		linkList = append(linkList, data.Link{
			ID:          link.ID,
			Url:         link.Url,
			Title:       link.Title,
			Description: link.Description,
			Keywords:    link.Keywords,
			StatusCode:  link.StatusCode,
			Remark:      link.Remark,
			Watching:    link.Watching,
		})
	}

	return linkList, nil
}
