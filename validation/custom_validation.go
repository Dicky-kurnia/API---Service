package validation

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func RegisterCustomValidations(validate *validator.Validate) {
	err := validate.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		_, err := time.Parse(fl.Param(), fl.Field().String())
		if err != nil {
			return false
		}
		return true
	})
	if err != nil {
		panic(err)
	}
}
