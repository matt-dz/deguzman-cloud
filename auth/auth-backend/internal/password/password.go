package password

import (
	"errors"
	"regexp"
)

var (
	uppercasePattern   = regexp.MustCompile(`[A-Z]`)
	numberPattern      = regexp.MustCompile(`\d`)
	specialCharPattern = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
)

var (
	ErrPasswordTooShort        = errors.New("password is too short")
	ErrPasswordNoUppercase     = errors.New("password must contain at least one uppercase letter")
	ErrPasswordNoNumber        = errors.New("password must contain at least one number")
	ErrorPasswordNoSpecialChar = errors.New("password must contain at least one special character")
)

func ValidatePassword(password string) error {
	if len(password) < 10 {
		return ErrPasswordTooShort
	}

	if !uppercasePattern.MatchString(password) {
		return ErrPasswordNoUppercase
	}

	if !numberPattern.MatchString(password) {
		return ErrPasswordNoNumber
	}

	if !specialCharPattern.MatchString(password) {
		return ErrorPasswordNoSpecialChar
	}

	return nil
}
