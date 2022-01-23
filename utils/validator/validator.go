package validator

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var (
	epsilon = math.Nextafter(1, 2) - 1
)

func init() {
	validate = validator.New()
	validate.RegisterValidation("is-monetary", isMonetaryValidate)
}

func Validate(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		vErrs, ok := err.(validator.ValidationErrors)
		if ok {
			errorStrings := []string{}
			for _, vErr := range vErrs {
				errorStrings = append(errorStrings, getErrorMessage(vErr))
			}

			return errors.New(strings.Join(errorStrings, "\n"))
		}

		return err
	}

	return nil
}

func isMonetaryValidate(fl validator.FieldLevel) bool {
	v := fl.Field().Float()
	return (v*1e2 - math.Floor(v*1e2)) < epsilon
}

func getErrorMessage(fErr validator.FieldError) string {
	switch fErr.Tag() {
	case "gt":
		return fmt.Sprintf("%s must be larger than %s", fErr.Field(), fErr.Param())
	case "gte":
		return fmt.Sprintf("%s must be larger or equal %s", fErr.Field(), fErr.Param())
	case "lte":
		return fmt.Sprintf("%s must be lesser or equal %s", fErr.Field(), fErr.Param())
	case "is-monetary":
		return fmt.Sprintf("%s is not in monetary", fErr.Field())
	}

	return fErr.Error()
}
