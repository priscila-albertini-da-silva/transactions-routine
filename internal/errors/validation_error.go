package errors

import "fmt"

const BadRequestErrorCode = 400

type ValidationError struct {
	Message string
	Code    int
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Validation error on: %s, code: %d", e.Message, e.Code)
}

func NewValidationError(msg string) *ValidationError {
	return &ValidationError{
		Message: msg,
		Code:    BadRequestErrorCode,
	}
}
