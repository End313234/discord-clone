package utils

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequestBody(body []byte, ctx *fiber.Ctx, v any) error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	})

	err := json.Unmarshal(body, v)
	if err != nil {
		return errors.New("invalid body")
	}

	if err = validate.Struct(v); err != nil {
		errs := err.(validator.ValidationErrors)
		parsedErrors := []string{}
		for _, e := range errs {
			if e.Param() != "" {
				parsedErrors = append(parsedErrors, ValidateError(e.Tag(), e.Field(), e.Param()))
			} else {
				parsedErrors = append(parsedErrors, ValidateError(e.Tag(), e.Field()))
			}
		}

		return errors.New(strings.Join(parsedErrors, ", "))
	}

	return err
}
