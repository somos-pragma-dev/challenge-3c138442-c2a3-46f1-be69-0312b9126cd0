package errors

import (
	"errors"
	"net/http"
)

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}