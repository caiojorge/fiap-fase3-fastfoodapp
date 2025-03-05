package validator

import "regexp"

type EmailValidator struct {
}

// IsValid valida o formato do string enviado no padrão email
func (v *EmailValidator) IsValid(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}
