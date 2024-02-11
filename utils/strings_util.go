package utils

import (
	"fmt"
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
