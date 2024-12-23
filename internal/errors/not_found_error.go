package errors

import "fmt"

const NotFoundErrorCode = 404

type NotFoundError struct {
	Message string
	Code    int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Not found error on: %s, code: %d", e.Message, e.Code)
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{
		Message: msg,
		Code:    BadRequestErrorCode,
	}
}
