package http

import (
	"comm/errors"
	"encoding/json"
	"net/http"
)

// WriteJSON defined TODO
func WriteJSON(w http.ResponseWriter, r *http.Request, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	rsp, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(rsp)
	if err != nil {
		return err
	}
	return nil
}

// WriteError defined TODO
func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	// parse out the error code
	ce := errors.Parse(err.Error())
	switch ce.Code {
	case 0:
		ce.Code = 500
		ce.Status = http.StatusText(500)
		ce.Detail = err.Error()
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ce.Error()))
}