package validations

import (
	"regexp"
	"unicode"
)

var Regex_correo = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidarPassword(s string) bool {
	var (
		hasMiLen  = false
		hasUpper  = false
		hasLover  = false
		hasNumber = false
	)

	if len(s) >= 6 && len(s) <= 20 {
		hasMiLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLover = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMiLen && hasUpper && hasLover && hasNumber
}
