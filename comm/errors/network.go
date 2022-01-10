package errors

import (
	"fmt"
	"net/http"

	"github.com/2637309949/micro/v3/service/errors"
)

const (
	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusPaymentRequired              = 402
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusMethodNotAllowed             = 405
	StatusNotAcceptable                = 406
	StatusProxyAuthRequired            = 407
	StatusRequestTimeout               = 408
	StatusConflict                     = 409
	StatusGone                         = 410
	StatusLengthRequired               = 411
	StatusPreconditionFailed           = 412
	StatusRequestEntityTooLarge        = 413
	StatusRequestURITooLong            = 414
	StatusUnsupportedMediaType         = 415
	StatusRequestedRangeNotSatisfiable = 416
	StatusExpectationFailed            = 417
	StatusTeapot                       = 418
	StatusMisdirectedRequest           = 421
	StatusUnprocessableEntity          = 422
	StatusLocked                       = 423
	StatusFailedDependency             = 424
	StatusTooEarly                     = 425
	StatusUpgradeRequired              = 426
	StatusPreconditionRequired         = 428
	StatusTooManyRequests              = 429
	StatusRequestHeaderFieldsTooLarge  = 431
	StatusUnavailableForLegalReasons   = 451

	StatusInternalServerError           = 500
	StatusNotImplemented                = 501
	StatusBadGateway                    = 502
	StatusServiceUnavailable            = 503
	StatusGatewayTimeout                = 504
	StatusHTTPVersionNotSupported       = 505
	StatusVariantAlsoNegotiates         = 506
	StatusInsufficientStorage           = 507
	StatusLoopDetected                  = 508
	StatusNotExtended                   = 510
	StatusNetworkAuthenticationRequired = 511
)

// BadRequest generates a 400 error.
func BadRequest(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusBadRequest,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusBadRequest),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusUnauthorized,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusUnauthorized),
	}
}

// Forbidden generates a 403 error.
func Forbidden(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusForbidden,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusForbidden),
	}
}

// NotFound generates a 404 error.
func NotFound(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusNotFound,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusNotFound),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusMethodNotAllowed,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusMethodNotAllowed),
	}
}

// Timeout generates a 408 error.
func Timeout(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusRequestTimeout,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusRequestTimeout),
	}
}

// Conflict generates a 409 error.
func Conflict(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusConflict,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusConflict),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusInternalServerError,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusInternalServerError),
	}
}

// NotImplemented generates a 501 error
func NotImplemented(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusNotImplemented,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusNotImplemented),
	}
}

// BadGateway generates a 502 error
func BadGateway(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusBadGateway,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusBadGateway),
	}
}

// ServiceUnavailable generates a 503 error
func ServiceUnavailable(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusServiceUnavailable,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusServiceUnavailable),
	}
}

// GatewayTimeout generates a 504 error
func GatewayTimeout(format string, a ...interface{}) error {
	return &errors.Error{
		Code:   StatusGatewayTimeout,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusGatewayTimeout),
	}
}
