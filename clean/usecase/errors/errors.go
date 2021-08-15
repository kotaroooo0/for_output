package errors

import (
	"fmt"
)

const (
	PostNotFoundCode string = "PostNotFound"
	UserNotFoundCode string = "UserNotFound"
)

type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code: %s message: %s", e.Code, e.Message)
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return &AppError{
		Message: err.Error(),
	}
}
