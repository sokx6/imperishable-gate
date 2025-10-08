package links

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/types/data"
	"strings"
)

// ListLinksByKeyword 根据关键词搜索用户的链接
// 支持在URL、备注、标题、描述、关键词、标签名和别名中进行模糊搜索，并返回分页结果
func ListLinksByKeyword(userID uint, keyword string, page int, pageSize int) ([]data.Link, error) {
	var links []model.Link

	// 1. 校验关键词，防止空值导致全表扫描
	if strings.TrimSpace(keyword) == "" {
		return []data.Link{}, nil
	}

	// 2. 构建模糊搜索模式
	searchPattern := "%" + keyword + "%"

	// 3. 多表联查并分页查询
	// 联接link_tags、tags和names表，在多个字段中搜索关键词
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

	// 4. 转换数据模型为响应格式
	var linkList []data.Link
	for _, link := range links {
		linkList = append(linkList, data.Link{
			Url:         link.Url,
			Tags:        utils.ExtractTagNames(link.Tags),
			Names:       utils.ExtractNames(link.Names),
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
