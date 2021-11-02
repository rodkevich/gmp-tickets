package validation

import (
	"fmt"
	"net/url"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

// Validate ...
func Validate(val *validator.Validate, model interface{}) []string {
	err := val.Struct(model)
	if err != nil {
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("the field \"%s\" is %s", err.Field(), err.Tag())
			errs = append(errs, errorMessage)
		}
		return errs
	}
	return nil
}

// IsValidLink ...
func IsValidLink(in string) bool {
	out, err := url.Parse(in)
	return err == nil && out.Scheme != "" && out.Host != ""
}
