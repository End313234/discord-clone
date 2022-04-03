package utils

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Constraints []string
	Error       string
}

var ValidationErrors map[string]string = map[string]string{
	"required": "%s is required",
	"email":    "%s must be a valid email address",
	"gte":      "%s must be greater than or equal to %s",
}

func ValidateError(tag string, constraints ...any) string {
	return fmt.Sprintf(ValidationErrors[tag], constraints...)
}

func RegisterTranslation(field string, text string, translator ut.Translator, validate *validator.Validate) {
	validate.RegisterTranslation(
		field,
		translator,
		func(universalTranslator ut.Translator) error {
			return universalTranslator.Add(field, text, true)
		},
		func(universalTranslator ut.Translator, fe validator.FieldError) string {
			t, _ := universalTranslator.T(field, fe.Field())
			return t
		},
	)
}

func OverrideTranslations(translator ut.Translator, validate *validator.Validate) {
	RegisterTranslation("required", "-{0} is required-", translator, validate)
	RegisterTranslation("email", "-{0} must be a valid email address-", translator, validate)
}
