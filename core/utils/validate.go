package utils

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(req interface{}) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}
