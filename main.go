package main

import (
	"Service-API/config"
	"Service-API/controller"
	"Service-API/exception"
	"Service-API/repository"
	"Service-API/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
}

func main() {
	heyvoMysql := config.MysqlHeyvoUtilitiesConnection()

	// Initialize Repositories Here
	adminRepository := repository.NewAdminRepository(heyvoMysql)

	// Initialize Services Here
	authService := service.NewAuthService(adminRepository)

	// Initialize Controllers Here
	authController := controller.NewAuthController(authService)

	// Fiber Setup Here
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New())

	// Router Setup Here
	v1 := app.Group("api/v1")
	authController.Route(v1)

	// Start App
	err := app.Listen(":" + os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}
