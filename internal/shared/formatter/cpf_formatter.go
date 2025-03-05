package formatter

import (
	"fmt"
	"regexp"
)

func RemoveMaskFromCPF(cpf string) string {
	re := regexp.MustCompile(`[^\d]`)
	return re.ReplaceAllString(cpf, "")
}

// PutMaskOnCPF formats a CPF string
func PutMaskOnCPF(cpf string) (string, error) {
	// Remove any non-digit characters
	cleanedCPF := RemoveMaskFromCPF(cpf)

	// Check if the cleaned CPF has exactly 11 digits
	if len(cleanedCPF) != 11 {
		return "", fmt.Errorf("invalid CPF length")
	}

	// Format the CPF string
	formattedCPF := fmt.Sprintf("%s.%s.%s-%s",
		cleanedCPF[:3],
		cleanedCPF[3:6],
		cleanedCPF[6:9],
		cleanedCPF[9:11],
	)

	return formattedCPF, nil
}
