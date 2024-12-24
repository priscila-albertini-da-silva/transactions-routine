package validator_test

import (
	"testing"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/errors"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/validator"
	"github.com/stretchr/testify/suite"
)

type AccountValidatorTestSuite struct {
	suite.Suite
}

func (suite *AccountValidatorTestSuite) TestValidate_ValidCPF() {
	// given: valid CPF
	account := buildValidAccount()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: there should be no error
	suite.NoError(err)
}

func (suite *AccountValidatorTestSuite) TestValidate_InvalidCPF() {
	// given: invalid CPF
	account := buildInvalidAccount()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: should return a validation error
	suite.Error(err)
	suite.IsType(&errors.ValidationError{}, err)
	suite.Equal("Validation error on: invalid CPF, code: 400", err.Error())
}

func (suite *AccountValidatorTestSuite) TestValidate_CPFWithNonNumericChars() {
	// given: CPF containing non-numeric characters
	account := buildValidAccount()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: there should be no error, as non-numeric characters are removed
	suite.NoError(err)
}

func (suite *AccountValidatorTestSuite) TestValidate_CPFWithIncorrectLength() {
	// given: CPF of incorrect length
	account := buildShortCPFAccount()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: should return a validation error
	suite.Error(err)
	suite.IsType(&errors.ValidationError{}, err)
	suite.Equal("Validation error on: invalid CPF, code: 400", err.Error())
}

func (suite *AccountValidatorTestSuite) TestValidate_CPFWithInvalidCheckDigits() {
	// given: CPF with invalid check digits
	account := buildAccountWithInvalidCheckDigits()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: should return a validation error
	suite.Error(err)
	suite.IsType(&errors.ValidationError{}, err)
	suite.Equal("Validation error on: invalid CPF, code: 400", err.Error())
}

func (suite *AccountValidatorTestSuite) TestValidate_CPFWithSameDigits() {
	// given: CPF with all digits the same
	account := buildAccountWithSameDigits()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: should return a validation error
	suite.Error(err)
	suite.IsType(&errors.ValidationError{}, err)
	suite.Equal("Validation error on: invalid CPF, code: 400", err.Error())
}

func (suite *AccountValidatorTestSuite) TestValidate_CPFWithAdditionalNonNumericChars() {
	// given: CPF containing additional non-numeric characters
	account := buildAccountWithSpecialChars()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: there should be no error, as non-numeric characters will be removed
	suite.NoError(err)
}

func (suite *AccountValidatorTestSuite) TestValidate_CPFWithSpaces() {
	// given: CPF containing spaces
	account := buildAccountWithSpaces()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: there should be no error, as spaces will be removed
	suite.NoError(err)
}

func (suite *AccountValidatorTestSuite) TestValidate_CPFWithSpecialCharacters() {
	// given: CPF containing special characters
	account := buildAccountWithSpecialChar()

	// when: tries to validate the CPF
	validator := validator.NewAccountValidator()
	err := validator.Validate(&account)

	// then: should return an error, as the CPF is not valid
	suite.Error(err)
	suite.IsType(&errors.ValidationError{}, err)
	suite.Equal("Validation error on: invalid CPF, code: 400", err.Error())
}

func TestAccountValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(AccountValidatorTestSuite))
}

func buildValidAccount() entity.Account {
	return entity.Account{
		DocumentNumber: "123.456.789-09",
	}
}

func buildInvalidAccount() entity.Account {
	return entity.Account{
		DocumentNumber: "123.456.789-00",
	}
}

func buildShortCPFAccount() entity.Account {
	return entity.Account{
		DocumentNumber: "1234567890",
	}
}

func buildAccountWithInvalidCheckDigits() entity.Account {
	return entity.Account{
		DocumentNumber: "123.456.789-00",
	}
}

func buildAccountWithSameDigits() entity.Account {
	return entity.Account{
		DocumentNumber: "111.111.111-11",
	}
}

func buildAccountWithSpecialChars() entity.Account {
	return entity.Account{
		DocumentNumber: "123@456$789#09",
	}
}

func buildAccountWithSpaces() entity.Account {
	return entity.Account{
		DocumentNumber: " 123 456 789 09 ",
	}
}

func buildAccountWithSpecialChar() entity.Account {
	return entity.Account{
		DocumentNumber: "123.456.789-º9",
	}
}
