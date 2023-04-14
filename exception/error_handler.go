package exception

import (
	"Service-API/exception/listerr"
	"Service-API/model"
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
)

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var customError *listerr.CustomErr
	if errors.As(err, &customError) {
		return ctx.Status(customError.Status()).JSON(model.Response{
			Code:   customError.Status(),
			Status: customError.Code(),
			Data:   nil,
			Error: map[string]interface{}{
				"general": customError.Error(),
			},
		})
	}

	return ctx.Status(500).JSON(model.Response{
		Code:   500,
		Status: INTERNAL_SERVER_ERROR,
		Data:   err.Error(),
	})
}
