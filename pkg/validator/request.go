package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strings"
)

type RequestValidator struct {
	Validator *validator.Validate
}

func (cv *RequestValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err == nil {
		return nil
	}

	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	errors := make(map[string]string)
	for _, fe := range err.(validator.ValidationErrors) {
		jsonField := getJSONFieldName(t, fe.StructField())
		errors[jsonField] = getErrorMessage(jsonField, fe)
	}

	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"errors": errors,
	})
}

// Get JSON tag from struct field
func getJSONFieldName(t reflect.Type, field string) string {
	if f, ok := t.FieldByName(field); ok {
		tag := f.Tag.Get("json")
		if tag != "" && tag != "-" {
			return strings.Split(tag, ",")[0]
		}
	}
	return strings.ToLower(field)
}

func getErrorMessage(field string, fe validator.FieldError) string {
	fieldTitle := strings.Title(field)

	switch fe.Tag() {
	case "required", "required_if", "required_unless", "required_with", "required_without":
		return fmt.Sprintf("%s is required", fieldTitle)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fieldTitle, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", fieldTitle, fe.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", fieldTitle, fe.Param())
	case "eq":
		return fmt.Sprintf("%s must be equal to %s", fieldTitle, fe.Param())
	case "ne":
		return fmt.Sprintf("%s must not be equal to %s", fieldTitle, fe.Param())
	case "eqfield", "eqcsfield":
		return fmt.Sprintf("%s must match %s", fieldTitle, fe.Param())
	case "nefield", "necsfield":
		return fmt.Sprintf("%s must not match %s", fieldTitle, fe.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", fieldTitle, fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", fieldTitle, fe.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", fieldTitle, fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", fieldTitle, fe.Param())
	case "gtfield", "gtefield", "ltfield", "ltefield":
		return fmt.Sprintf("%s must compare correctly with %s", fieldTitle, fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldTitle)
	case "url", "uri", "http_url":
		return fmt.Sprintf("%s must be a valid URL", fieldTitle)
	case "uuid", "uuid4", "uuid5":
		return fmt.Sprintf("%s must be a valid UUID", fieldTitle)
	case "datetime":
		return fmt.Sprintf("%s must be a valid datetime", fieldTitle)
	case "alpha":
		return fmt.Sprintf("%s must contain only letters", fieldTitle)
	case "alphanum":
		return fmt.Sprintf("%s must contain only letters and numbers", fieldTitle)
	case "numeric":
		return fmt.Sprintf("%s must be numeric", fieldTitle)
	case "boolean":
		return fmt.Sprintf("%s must be true or false", fieldTitle)
	case "ip", "ipv4", "ipv6":
		return fmt.Sprintf("%s must be a valid IP address", fieldTitle)
	case "hostname":
		return fmt.Sprintf("%s must be a valid hostname", fieldTitle)
	case "contains":
		return fmt.Sprintf("%s must contain %s", fieldTitle, fe.Param())
	case "excludes":
		return fmt.Sprintf("%s must not contain %s", fieldTitle, fe.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", fieldTitle, fe.Param())
	case "file":
		return fmt.Sprintf("%s must be a valid file", fieldTitle)
	case "dir":
		return fmt.Sprintf("%s must be a valid directory", fieldTitle)
	case "base64", "json":
		return fmt.Sprintf("%s must be a valid %s", fieldTitle, fe.Tag())
	case "unique":
		return fmt.Sprintf("%s must be unique", fieldTitle)
	default:
		return fmt.Sprintf("%s is invalid", fieldTitle)
	}
}
