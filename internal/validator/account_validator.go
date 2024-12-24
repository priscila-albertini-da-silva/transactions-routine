package validator

import (
	"regexp"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/errors"
	"go.uber.org/fx"
)

type AccountValidator interface {
	Validate(account *entity.Account) error
}

type AccountValidatorImpl struct{}

func NewAccountValidator() AccountValidator {
	return &AccountValidatorImpl{}
}

func (v *AccountValidatorImpl) Validate(account *entity.Account) error {
	if !v.isValidCPF(account.DocumentNumber) {
		return errors.NewValidationError("invalid CPF")
	}
	return nil
}

func (v *AccountValidatorImpl) isValidCPF(cpf string) bool {
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
	return byte(cpf[9]) == d1 && byte(cpf[10]) == d2
}

func calculateCPFCheckDigits(cpfPrefix string) (byte, byte) {
	// Primeiro dígito verificador
	sum := 0
	for i, c := range cpfPrefix {
		sum += int(c-'0') * (10 - i)
	}
	firstDigit := 11 - (sum % 11)
	if firstDigit >= 10 {
		firstDigit = 0
	}

	// Adiciona o primeiro dígito para o cálculo do segundo
	cpfPrefix += string(firstDigit + '0')

	// Segundo dígito verificador
	sum = 0
	for i, c := range cpfPrefix {
		sum += int(c-'0') * (11 - i)
	}
	secondDigit := 11 - (sum % 11)
	if secondDigit >= 10 {
		secondDigit = 0
	}

	return byte(firstDigit + '0'), byte(secondDigit + '0')
}

var ModuleAccountValidator = fx.Module("account_validator", fx.Invoke(NewAccountValidator))
