package remarks

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/service/common"
)

func AddRemarkByName(name string, userId uint, remark string) error {
	var id uint
	if id = utils.GetLinkIDByName(name, userId); id == 0 {
		return common.ErrNameNotFound
	} else if err := database.DB.Model(&model.Link{ID: id}).Update("Remark", remark).Error; err != nil {
		return common.ErrDatabase
	}
	return nil
}
