package utils

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	var str = ""
	for i := 0; i < len(s); i++ {
		if i == 0 && unicode.IsUpper(rune(s[i])) {
			str = strings.ToLower(string(s[i]))
			continue
		}
		if i != 0 && unicode.IsUpper(rune(s[i])) {
			str += fmt.Sprintf("_%s", strings.ToLower(string(s[i])))
		} else {
			str += string(s[i])
		}
	}
	return str
}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regExp := regexp.MustCompile(pattern)

	return regExp.MatchString(email)
}
