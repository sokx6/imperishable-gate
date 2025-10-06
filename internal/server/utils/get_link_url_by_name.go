package utils

//通过查询数据库将name转换为linkId
import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func GetLinkUrlByName(name string, userId uint) string {
	linkId := GetLinkIDByName(name, userId)
	if linkId == 0 {
		return ""
	}
	var link model.Link
	if err := database.DB.Where("id = ? AND user_id = ?", linkId, userId).Take(&link).Error; err != nil {
		return ""
	}
	return link.Url
}
