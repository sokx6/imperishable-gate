package utils

//通过查询数据库将name转换为linkId
import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func NameToLinkId(name string, userId uint) uint {
	var Name model.Name
	if err := database.DB.Where("name = ? AND user_id = ?", name, userId).Take(&Name).Error; err != nil {
		return 0
	}
	return Name.LinkID
}
