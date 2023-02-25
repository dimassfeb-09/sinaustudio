package helpers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ErrorValidateHandler(err error) string {
	var Msg string
	switch err.(type) {
	case validator.ValidationErrors:
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			tag := err.Tag()
			Msg = fmt.Sprintf("Key: %s, Tag: %s", field, tag)
			break
		}
	case *json.UnmarshalTypeError:
		errUnmarshal := err.(*json.UnmarshalTypeError)
		field := errUnmarshal.Field
		tagType := errUnmarshal.Type
		Msg = fmt.Sprintf("Key: %s, tipe data yang diizinkan: %s.", field, tagType)
	default:
		Msg = err.Error()
	}

	return Msg

}
