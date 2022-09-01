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
func BadRequest(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusBadRequest,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusBadRequest),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusUnauthorized,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusUnauthorized),
	}
}

// Forbidden generates a 403 error.
func Forbidden(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusForbidden,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusForbidden),
	}
}

// NotFound generates a 404 error.
func NotFound(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusNotFound,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusNotFound),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusMethodNotAllowed,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusMethodNotAllowed),
	}
}

// Timeout generates a 408 error.
func Timeout(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusRequestTimeout,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusRequestTimeout),
	}
}

// Conflict generates a 409 error.
func Conflict(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusConflict,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusConflict),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusInternalServerError,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusInternalServerError),
	}
}

// NotImplemented generates a 501 error
func NotImplemented(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusNotImplemented,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusNotImplemented),
	}
}

// BadGateway generates a 502 error
func BadGateway(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusBadGateway,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusBadGateway),
	}
}

// ServiceUnavailable generates a 503 error
func ServiceUnavailable(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusServiceUnavailable,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusServiceUnavailable),
	}
}

// GatewayTimeout generates a 504 error
func GatewayTimeout(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusGatewayTimeout,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(StatusGatewayTimeout),
	}
}
