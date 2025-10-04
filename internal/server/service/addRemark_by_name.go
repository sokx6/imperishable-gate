package service

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
)

func AddRemarkByName(name string, remark string) error {
	var Name model.Name
	if err := database.DB.Where("name = ?", name).Take(&Name).Error; err != nil {
		return ErrNameNotFound
	}
	if err := database.DB.Model(&model.Link{ID: Name.LinkID}).Update("Remark", remark).Error; err != nil {
		return ErrDatabase
	}
	return nil
}
