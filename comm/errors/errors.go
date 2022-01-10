package errors

import (
	"encoding/json"
	"net/http"

	"github.com/2637309949/micro/v3/service/errors"
)

// New generates a custom error.
func New(detail string, code int32) error {
	return &errors.Error{
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// FromError try to convert go error to *Error
func FromError(err error) *errors.Error {
	if err == nil {
		return nil
	}
	if verr, ok := err.(*errors.Error); ok && verr != nil {
		return verr
	}

	return Parse(err.Error())
}

// Parse tries to parse a JSON string into an error. If that
// fails, it will set the given string as the error detail.
func Parse(err string, trace ...string) *errors.Error {
	e := new(errors.Error)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		e.Detail = err
	}
	if len(trace) > 0 && e.RequestId == "" {
		e.RequestId = trace[0]
	}
	return e
}

func Is(err error, errf func(string, ...interface{}) error) bool {
	sErr, tErr := FromError(errf(err.Error())), FromError(err)
	return sErr.Code == tErr.Code
}
