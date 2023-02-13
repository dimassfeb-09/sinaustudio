package handlers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ErrorValidateHandler(err error) []string {
	var errorList []string
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		tag := err.Tag()
		errorList = append(errorList, "Field with key "+field+" are "+tag)
	}
	return errorList
}
