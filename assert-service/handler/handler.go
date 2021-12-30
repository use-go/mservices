package handler

import (
	"net/http"

	"comm/errors"
)

type Handler struct{}

func writeError(w http.ResponseWriter, err error) {
	// response content type
	w.Header().Set("Content-Type", "application/json")

	// parse out the error code
	ce := errors.Parse(err.Error())

	switch ce.Code {
	case 0:
		ce.Code = 500
		ce.Status = http.StatusText(500)
		ce.Detail = "error during request: " + ce.Detail
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusOK)
	}

	w.Write([]byte(ce.Error()))
}
