package errors

import (
	"net/http"

	"github.com/2637309949/micro/v3/service/errors"
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

// FromError try to convert go error to *Error
func FromError(err error) *errors.Error {
	return errors.FromError(err)
}

// Parse tries to parse a JSON string into an error. If that
func Parse(err string, trace ...string) *errors.Error {
	return errors.Parse(err, trace...)
}

func Is(err error, errf func(string, ...interface{}) error) bool {
	sErr, tErr := FromError(errf(err.Error())), FromError(err)
	return sErr.Code == tErr.Code
}
