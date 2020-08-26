package validators

import (
	"strings"

	"github.com/go-playground/validator"
)

//TitleValidate function to validate the cool validation
func TitleValidate(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "cool")
}
