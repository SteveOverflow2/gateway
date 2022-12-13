package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	ErrorCodeNotFound     = "not_found"
	ErrorCodeInvalid      = "invalid"
	ErrorCodeInternal     = "internal"
	ErrorCodeUnauthorized = "unauthorized"
	ErrorCodeConflict     = "conflict"
	ErrorCodeForbidden    = "forbidden"
)

// This struct represents the internal error struct.
type Error struct {
	Origin  error
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) ErrorMessage() string {
	return e.Message
}

func (e *Error) ErrorCode() string {
	return e.Code
}

func (e *Error) ErrorOrigin() error {
	return e.Origin
}

const errorLog = "[Error]: "

// // WrapErrorf returns a wrapped error.
func WrapErrorf(origin error, code string, format string, args ...interface{}) error {
	log.Println(errorLog + origin.Error())

	return &Error{
		Origin:  origin,
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// NewErrorf instantiates a new error.
func NewErrorf(origin error, code string, format string, args ...interface{}) error {
	return WrapErrorf(origin, code, format, args...)
}

func GetHTTPCode(err error) int {
	var internalErr *Error
	if !errors.As(err, &internalErr) {
		return http.StatusInternalServerError
	}

	switch internalErr.Code {
	case ErrorCodeNotFound:
		return http.StatusNotFound
	case ErrorCodeInvalid:
		return http.StatusBadRequest
	case ErrorCodeInternal:
		return http.StatusInternalServerError
	case ErrorCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrorCodeConflict:
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}
