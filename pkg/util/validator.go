package util

import (
	"strings"

	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/go-playground/validator/v10"
)

func Validate(v *validator.Validate, req any) (map[string][]string, error) {
	err := v.Struct(req)
	messages := make(map[string][]string)
	if err != nil {
		trans := config.GetTranslator()
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			translated := strings.ToLower(e.Translate(trans))
			messages[field] = append(messages[field], translated)
		}
	}
	return messages, err
}
