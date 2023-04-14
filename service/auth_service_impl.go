package service

import (
	"Service-API/exception/listerr"
	"Service-API/helper"
	"Service-API/model"
	"Service-API/repository"
	"Service-API/validation"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	AdminRepository repository.AdminRepository
}

func (service authService) Login(request model.LoginRequest) (string, error) {
	err := validation.Validate(request)
	if err != nil {
		return "", err
	}

	admin, err := service.AdminRepository.GetAdminByUsername(request.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", listerr.USERNAME_OR_PASSWORD_INVALID
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password))
	if err != nil {
		return "", listerr.USERNAME_OR_PASSWORD_INVALID
	}

	token, err := helper.EncodeToken(admin.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service authService) Logout(tokenString string) error {
	_, err := helper.ValidateToken(tokenString)
	if err != nil {
		return listerr.UNAUTHORIZED
	}

	decoded, err := helper.DecodeToken(tokenString)
	if err != nil {
		return listerr.UNAUTHORIZED
	}

	helper.DelRedis(fmt.Sprintf("cmsv1-token-%s", decoded.AccessUUID))
	return nil
}

func NewAuthService(adminRepository repository.AdminRepository) AuthService {
	return &authService{AdminRepository: adminRepository}
}
