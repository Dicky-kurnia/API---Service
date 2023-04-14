package repository

import "Service-API/entity"

type AdminRepository interface {
	GetAdminByUsername(username string) (entity.AdminEntity, error)
}
