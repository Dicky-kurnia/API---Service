package controller

import (
	"Service-API/exception"
	"Service-API/middleware"
	"Service-API/model"
	"Service-API/service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type AuthController struct {
	AuthService service.AuthService
}

func (controller AuthController) Route(app fiber.Router) {
	app = app.Group("auth")

	app.Post("", controller.Login)
	app.Post("/logout", middleware.CheckToken(), controller.Logout)
}

func (controller AuthController) Login(c *fiber.Ctx) error {
	request := new(model.LoginRequest)
	err := c.BodyParser(request)
	exception.PanicIfNeeded(err)

	token, err := controller.AuthService.Login(*request)
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data: map[string]interface{}{
			"access_token": token,
		},
	})
}

func (controller AuthController) Logout(c *fiber.Ctx) error {
	tokenSlice := strings.Split(c.Get("Authorization"), "Bearer ")
	var tokenString string
	if len(tokenSlice) == 2 {
		tokenString = tokenSlice[1]
	}

	err := controller.AuthService.Logout(tokenString)
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
	})
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}
