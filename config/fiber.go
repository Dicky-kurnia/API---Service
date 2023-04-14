package config

import (
	"Service-API/exception"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: exception.ErrorHandler,
		BodyLimit:    5 * 1024 * 1024,
	}
}
