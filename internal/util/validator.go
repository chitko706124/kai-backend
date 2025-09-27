package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	validate := validator.New()
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.StructField()
		tag := err.Tag()
		param := err.Param()

		switch tag {
		case "required":
			errors[field] = fmt.Sprintf("Field %s is required", field)
		case "email":
			errors[field] = fmt.Sprintf("Field %s must be a valid email address", field)
		case "min":
			errors[field] = fmt.Sprintf("Field %s must be at least %s characters long", field, param)
		case "max":
			errors[field] = fmt.Sprintf("Field %s must be at most %s characters long", field, param)
		case "eqfield":
			errors[field] = fmt.Sprintf("Field %s must be the same as %s", field, param)
		default:
			errors[field] = err.Error()
		}
	}
	return errors
}
