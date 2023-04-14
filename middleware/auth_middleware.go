package middleware

import (
	"Service-API/exception/listerr"
	"Service-API/helper"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CheckToken() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenSlice := strings.Split(c.Get("Authorization"), "Bearer ")
		var tokenString string
		if len(tokenSlice) == 2 {
			tokenString = tokenSlice[1]
		}

		// validate token
		_, err := helper.ValidateToken(tokenString)
		if err != nil {
			fmt.Println(err)
			return listerr.UNAUTHORIZED
		}

		// extract data from token
		decodedRes, err := helper.DecodeToken(tokenString)
		if err != nil {
			fmt.Println(err)
			return listerr.UNAUTHORIZED
		}

		ok, dataRedis := helper.GetRedis[string](fmt.Sprintf("cmsv1-token-%s", decodedRes.AccessUUID))
		if !ok || dataRedis == "" {
			return listerr.UNAUTHORIZED
		}

		// set to global var
		c.Locals("currentAdminId", decodedRes.AdminId)
		return c.Next()
	}
}
