package errors

import (
	"fmt"
	"net/http"

	"github.com/micro/micro/v3/service/errors"
)

// New generates a custom error.
func New(id, detail string, code int32) error {
	return &errors.Error{
		Id:     id,
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   500,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(500),
	}
}
