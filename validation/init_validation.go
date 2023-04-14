package validation

import (
	"Service-API/exception"
	"Service-API/exception/listerr"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func NewValidator() *validator.Validate {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Register the custom validation functions
	RegisterCustomValidations(validate)

	return validate
}

func Validate[T any](val T) error {
	validate := NewValidator()
	err := validate.Struct(val)
	if err != nil {
		var errorx exception.ValidationError
		for _, errors := range err.(validator.ValidationErrors) {
			if errorx.Message == "" {
				errorx = exception.ValidationError{
					Message: fmt.Sprintf("{\"%s\": \"%v\"}", errors.Field(), FormatValidationError(errors.Tag(), errors.Param(), errors.Value())),
				}
			} else {
				errorx = exception.ValidationError{
					Message: errorx.Message[:len(errorx.Message)-1] + "," + fmt.Sprintf("\"%s\": \"%v\"}", errors.Field(), FormatValidationError(errors.Tag(), errors.Param(), errors.Value())),
				}
			}

		}
		return errorx
	}
	return nil
}

func FormatValidationError(tag string, param interface{}, value interface{}) string {
	switch tag {
	case "required":
		return listerr.NOT_BLANK
	case "min":
		return listerr.Min(param)
	case "max":
		return listerr.Max(param)
	default:
		return listerr.NOT_VALID
	}
}
