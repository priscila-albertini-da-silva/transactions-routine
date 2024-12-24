package errors_test

import (
	"testing"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/errors"
	"github.com/stretchr/testify/suite"
)

type ValidationErrorTestSuite struct {
	suite.Suite
}

func (suite *ValidationErrorTestSuite) TestNewValidationError() {
	// given: a validation error message
	message := "Invalid input"

	// when: creating a new validation error
	validationError := errors.NewValidationError(message)

	// then: assert the expected result
	suite.Equal(message, validationError.Message)
	suite.Equal(errors.BadRequestErrorCode, validationError.Code)
}

func (suite *ValidationErrorTestSuite) TestValidationErrorErrorMethod() {
	// given: a validation error with specific message and code
	message := "Invalid input"
	validationError := errors.NewValidationError(message)

	// when: calling the Error method
	result := validationError.Error()

	// then: assert the expected result
	expectedMessage := "Validation error on: Invalid input, code: 400"
	suite.Equal(expectedMessage, result)
}

func TestValidationErrorSuite(t *testing.T) {
	suite.Run(t, new(ValidationErrorTestSuite))
}
