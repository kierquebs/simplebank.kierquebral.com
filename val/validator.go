package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidusername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidName     = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func ValidateString(value string, minLenght int, maxLength int) error {
	n := len(value)
	if n < minLenght || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLenght, maxLength)
	}
	return nil

}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidusername(value) {
		return fmt.Errorf("must contains only lowercase letters, digits,  or underscore")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

func ValidateName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("must contains only letters or spaces")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
