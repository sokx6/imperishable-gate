package service

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
)

func AddRemarkByName(name string, userId uint, remark string) error {
	var id uint
	if id = utils.NameToLinkId(name, userId); id == 0 {
		return ErrNameNotFound
	} else if err := database.DB.Model(&model.Link{ID: id}).Update("Remark", remark).Error; err != nil {
		return ErrDatabase
	}
	return nil
}
