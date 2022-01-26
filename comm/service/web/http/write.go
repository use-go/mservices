package http

import (
	"bytes"
	"comm/errors"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/2637309949/micro/v3/service/debug/trace"
)

func OutputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}

// Success defined TODO
func Success(w http.ResponseWriter, r *http.Request, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	traceID, _, _ := trace.FromContext(r.Context())
	var err error
	var rsp []byte
	switch o := obj.(type) {
	case string:
		rsp = []byte(o)
	case []byte:
		rsp = o
	default:
		rsp, err = json.Marshal(obj)
		if err != nil {
			return err
		}
	}
	w.WriteHeader(http.StatusOK)
	rsp = bytes.TrimRight(rsp, "\n")
	if !strings.HasPrefix(string(rsp), "{") {
		switch obj.(type) {
		case string, []byte:
			rsp = []byte(fmt.Sprintf("{\"data\": %v}", strconv.Quote(string(rsp))))
		default:
			rsp = []byte(fmt.Sprintf(`{"data": %v}`, string(rsp)))
		}
	}
	if !strings.Contains(string(rsp), "code") {
		rsp = []byte(strings.Replace(string(rsp), "{", "{\"code\": 200,", 1))
	}
	if !strings.Contains(string(rsp), "request_id") {
		rsp = []byte(strings.Replace(string(rsp), "{", "{\"request_id\": \""+traceID+"\",", 1))
	}
	rsp = []byte(strings.ReplaceAll(string(rsp), ",}", "}"))
	_, err = w.Write(rsp)
	if err != nil {
		return err
	}
	return nil
}

// Fail defined TODO
func Fail(w http.ResponseWriter, r *http.Request, err error) {
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
