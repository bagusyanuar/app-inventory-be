package config

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var trans ut.Translator

func NewValidator() *validator.Validate {
	v := validator.New()
	locale := en.New()
	uni := ut.New(locale, locale)
	t, _ := uni.GetTranslator("en")
	trans = t

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Tag.Get("json")
		if tag == "-" || tag == "" {
			return fld.Name
		}
		return strings.Split(tag, ",")[0]
	})

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		panic("failed to register default EN translations: " + err.Error())
	}

	if err := RegisterSymbolValidation(v); err != nil {
		panic("failed to register symbol validation: " + err.Error())
	}

	if err := RegisterArrayValidation(v); err != nil {
		panic("failed to register array validation: " + err.Error())
	}

	if err := RegisterSymbolTranslation(v, trans); err != nil {
		panic("failed to register symbol translation: " + err.Error())
	}

	if err := RegisterArrayTranslation(v, trans); err != nil {
		panic("failed to register array translation: " + err.Error())
	}

	return v
}

func RegisterSymbolValidation(v *validator.Validate) error {
	return v.RegisterValidation("symbol", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		symbolRegex := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`)
		return symbolRegex.MatchString(value)
	})
}

func RegisterArrayValidation(v *validator.Validate) error {
	return v.RegisterValidation("array", func(fl validator.FieldLevel) bool {
		raw, ok := fl.Field().Interface().(json.RawMessage)
		if !ok {
			return false
		}

		var tmp any
		if err := json.Unmarshal(raw, &tmp); err != nil {
			return false
		}

		_, isArray := tmp.([]any) // cek apakah hasil decode adalah array JSON
		return isArray
	})
}

func RegisterSymbolTranslation(v *validator.Validate, trans ut.Translator) error {
	return v.RegisterTranslation("symbol", trans,
		func(ut ut.Translator) error {
			return ut.Add("symbol", "{0} must contain at least one symbol", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("symbol", fe.Field())
			return t
		},
	)
}

func RegisterArrayTranslation(v *validator.Validate, trans ut.Translator) error {
	return v.RegisterTranslation("array", trans,
		func(ut ut.Translator) error {
			return ut.Add("array", "{0} must be an array", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("array", fe.Field())
			return t
		},
	)
}

func GetTranslator() ut.Translator {
	return trans
}
