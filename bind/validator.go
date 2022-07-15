package bind

import (
	"strconv"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func Password(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	l, _ := strconv.Atoi(fl.Param())
	if l < 3 {
		l = 3
	}

	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)

	if len(s) >= l {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}

func PasswordSpecial(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	l, _ := strconv.Atoi(fl.Param())

	if l < 4 {
		l = 4
	}

	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= l {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
