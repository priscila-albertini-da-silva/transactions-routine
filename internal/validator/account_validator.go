package validator

import (
	"regexp"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/errors"
)

type AccountValidator struct{}

func NewAccountValidator() *AccountValidator {
	return &AccountValidator{}
}

func (v *AccountValidator) Validate(account *entity.Account) error {
	if !v.isValidCPF(account.DocumentNumber) {
		return errors.NewValidationError("invalid CPF")
	}
	return nil
}

func (v *AccountValidator) isValidCPF(cpf string) bool {
	cleanCPF := removeNonDigits(cpf)

	if len(cleanCPF) != 11 {
		return false
	}

	return isCPFValid(cleanCPF)
}

func removeNonDigits(cpf string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(cpf, "")
}

func isCPFValid(cpf string) bool {
	d1, d2 := calculateCPFCheckDigits(cpf[:9])
	return cpf[9] == d1 && cpf[10] == d2
}

func calculateCPFCheckDigits(cpfPrefix string) (byte, byte) {
	var sum int
	for i, c := range cpfPrefix {
		sum += int(c-'0') * (10 - i)
	}
	firstDigit := (sum * 10) % 11
	if firstDigit == 10 {
		firstDigit = 0
	}

	sum = 0
	for i, c := range cpfPrefix {
		sum += int(c-'0') * (11 - i)
	}
	secondDigit := (sum * 10) % 11
	if secondDigit == 10 {
		secondDigit = 0
	}

	return byte(firstDigit + '0'), byte(secondDigit + '0')
}
