package errors

import (
	"fmt"
	"net/http"

	"github.com/micro/micro/v3/service/errors"
)

// New generates a custom error.
func New(detail string, code int32) error {
	return &errors.Error{
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   500,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(500),
	}
}
