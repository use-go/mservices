package web

import (
	"bytes"
	"comm/trace"
	"io/ioutil"
	"net/http"

	cx "github.com/2637309949/micro/v3/util/ctx"
)

type responseBodyWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	// reduce mem
	if r.ResponseWriter.Header().Get("Content-Type") == "application/json" {
		r.body.Write(b)
	}
	return r.ResponseWriter.Write(b)
}

var (
	debugWrapper = func(call func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(res http.ResponseWriter, req *http.Request) {
			req = req.WithContext(cx.FromRequest(req))
			bodyBytes := []byte("{}")
			if req.Header.Get("Content-Type") == "application/json" {
				bodyBytes, _ := ioutil.ReadAll(req.Body)
				req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}
			w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: res}
			defer func() {
				trace.Debug(cx.FromRequest(req), req.URL.Path, bodyBytes, w.body.Bytes())()
			}()
			call(w, req)
		}
	}
)
