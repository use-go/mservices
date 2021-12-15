package errors

import (
	"fmt"
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

// BadRequest generates a 400 error.
func BadRequest(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   400,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(400),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   401,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusUnauthorized),
	}
}

// Forbidden generates a 403 error.
func Forbidden(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   403,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusForbidden),
	}
}

// NotFound generates a 404 error.
func NotFound(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   404,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusNotFound),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   405,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusMethodNotAllowed),
	}
}

// Timeout generates a 408 error.
func Timeout(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   408,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusRequestTimeout),
	}
}

// Conflict generates a 409 error.
func Conflict(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   409,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusConflict),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   500,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusInternalServerError),
	}
}

// NotImplemented generates a 501 error
func NotImplemented(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   501,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusNotImplemented),
	}
}

// BadGateway generates a 502 error
func BadGateway(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   502,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusBadGateway),
	}
}

// ServiceUnavailable generates a 503 error
func ServiceUnavailable(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   503,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusServiceUnavailable),
	}
}

// GatewayTimeout generates a 504 error
func GatewayTimeout(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   504,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusGatewayTimeout),
	}
}
