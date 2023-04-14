package service

import "Service-API/model"

type AuthService interface {
	Login(request model.LoginRequest) (string, error)
	Logout(tokenString string) error
}
