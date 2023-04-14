package repository

import (
	"Service-API/entity"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func (repository adminRepository) GetAdminByUsername(username string) (entity.AdminEntity, error) {
	var result entity.AdminEntity
	if err := repository.DB.Model(&result).Where("username = ?", username).Scan(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func NewAdminRepository(DB *gorm.DB) AdminRepository {
	return &adminRepository{DB: DB}
}
