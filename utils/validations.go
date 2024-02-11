package utils

import (
	"fmt"
	"regexp"
)

func ValidateStringField(fieldName string, input map[string]any, errors map[string][]string) *string {
	field, ok := input[fieldName]

	if !ok {
		return nil
	}

	fieldStr, ok := field.(string)
	if !ok && len(errors[fieldName]) == 0 {
		errors[fieldName] = make([]string, 0)
	}
	if !ok {
		errors[fieldName] = append(errors[fieldName], fmt.Sprintf("%s must be a string", fieldName))
	}

	return &fieldStr
}

func ValidateRequiredStringField(fieldName string, input map[string]any, errors map[string][]string) *string {

	_, ok := input[fieldName]
	if !ok && len(errors[fieldName]) == 0 {
		errors[fieldName] = make([]string, 0)
	}
	if !ok {
		errors[fieldName] = append(errors[fieldName], fmt.Sprintf("%s is required", fieldName))
		return nil
	}

	return ValidateStringField(fieldName, input, errors)
}

func ValidateUrlField(fieldName string, input map[string]any, errors map[string][]string) *string {
	field := ValidateStringField(fieldName, input, errors)

	if field == nil {
		return nil
	}

	ok := ValidateUrl(*field)
	if !ok && len(errors[fieldName]) == 0 {
		errors[fieldName] = make([]string, 0)
	}
	if !ok {
		errors[fieldName] = append(errors[fieldName], fmt.Sprintf("%s must be a valid url", fieldName))
	}
	return field

}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regExp := regexp.MustCompile(pattern)

	return regExp.MatchString(email)
}

func ValidateUrl(url string) bool {
	pattern := `^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`
	regExp := regexp.MustCompile(pattern)

	return regExp.MatchString(url)
}
