package utils

import (
	"fmt"
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
